package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

const (
	federationURL   = "https://signin.aws.amazon.com/federation"
	consoleURL      = "https://console.aws.amazon.com/"
	sessionDuration = 43200 // 12 hours (max for federation)
)

var profile string

func main() {
	rootCmd := &cobra.Command{
		Use:   "aws-console",
		Short: "Open the AWS Console in your browser using current credentials",
		Long: `Authenticates using your AWS credentials and opens the AWS Management Console
in your default web browser. If credentials are expired or missing, it will
attempt to run 'aws sso login' to refresh them.`,
		SilenceUsage: true,
		RunE:         run,
	}

	rootCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile to use (defaults to AWS_PROFILE env var)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	// Determine profile: flag takes precedence, then env var
	if profile == "" {
		profile = os.Getenv("AWS_PROFILE")
	}

	// Load AWS config
	var opts []func(*config.LoadOptions) error
	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Validate credentials with STS GetCallerIdentity
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Credentials are not valid, attempting SSO login...\n")
		if loginErr := ssoLogin(profile); loginErr != nil {
			return fmt.Errorf("SSO login failed: %w", loginErr)
		}

		// Reload config after SSO login
		cfg, err = config.LoadDefaultConfig(ctx, opts...)
		if err != nil {
			return fmt.Errorf("failed to load AWS config after SSO login: %w", err)
		}

		stsClient = sts.NewFromConfig(cfg)
		identity, err = stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {
			return fmt.Errorf("credentials still invalid after SSO login: %w", err)
		}
	}

	fmt.Printf("Authenticated as: %s\n", aws.ToString(identity.Arn))

	// Retrieve the resolved credentials
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve credentials: %w", err)
	}

	// If no session token (e.g. long-lived IAM user keys), request temporary credentials
	if creds.SessionToken == "" {
		fmt.Println("No session token found, requesting temporary credentials...")
		out, err := stsClient.GetSessionToken(ctx, &sts.GetSessionTokenInput{
			DurationSeconds: aws.Int32(sessionDuration),
		})
		if err != nil {
			return fmt.Errorf("failed to get temporary credentials: %w", err)
		}
		creds = aws.Credentials{
			AccessKeyID:     aws.ToString(out.Credentials.AccessKeyId),
			SecretAccessKey: aws.ToString(out.Credentials.SecretAccessKey),
			SessionToken:    aws.ToString(out.Credentials.SessionToken),
		}
	}

	// Build the federated console sign-in URL
	loginURL, err := buildConsoleURL(creds)
	if err != nil {
		return fmt.Errorf("failed to build console URL: %w", err)
	}

	fmt.Println("Opening AWS Console in your browser...")
	return openBrowser(loginURL)
}

// ssoLogin shells out to the AWS CLI to perform an SSO login.
func ssoLogin(profile string) error {
	args := []string{"sso", "login"}
	if profile != "" {
		args = append(args, "--profile", profile)
	}

	cmd := exec.Command("aws", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// buildConsoleURL uses the AWS federation endpoint to create a pre-authenticated
// console sign-in URL from temporary credentials.
func buildConsoleURL(creds aws.Credentials) (string, error) {
	// Marshal the session credentials as JSON
	sessionData := map[string]string{
		"sessionId":    creds.AccessKeyID,
		"sessionKey":   creds.SecretAccessKey,
		"sessionToken": creds.SessionToken,
	}

	sessionJSON, err := json.Marshal(sessionData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session: %w", err)
	}

	// Request a sign-in token from the federation endpoint
	tokenURL := fmt.Sprintf(
		"%s?Action=getSigninToken&SessionDuration=%d&Session=%s",
		federationURL,
		sessionDuration,
		url.QueryEscape(string(sessionJSON)),
	)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(tokenURL)
	if err != nil {
		return "", fmt.Errorf("failed to request signin token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read federation response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("federation endpoint returned HTTP %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp struct {
		SigninToken string `json:"SigninToken"`
	}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse signin token response: %w", err)
	}

	if tokenResp.SigninToken == "" {
		return "", fmt.Errorf("received empty signin token from federation endpoint")
	}

	// Construct the final console login URL
	loginURL := fmt.Sprintf(
		"%s?Action=login&Issuer=aws-console-cli&Destination=%s&SigninToken=%s",
		federationURL,
		url.QueryEscape(consoleURL),
		url.QueryEscape(tokenResp.SigninToken),
	)

	return loginURL, nil
}

// openBrowser opens the given URL in the user's default browser.
func openBrowser(targetURL string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler"}
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	args = append(args, targetURL)
	return exec.Command(cmd, args...).Start()
}
