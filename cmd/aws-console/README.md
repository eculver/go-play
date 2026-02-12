# aws-console

Opens the AWS Management Console in your default browser using your current AWS credentials -- no web login form required.

## How it works

1. Resolves your AWS profile from the `-p`/`--profile` flag or the `AWS_PROFILE` environment variable.
2. Validates credentials by calling STS `GetCallerIdentity`.
3. If credentials are expired or missing, automatically runs `aws sso login` to refresh them.
4. If the credentials are long-lived IAM keys (no session token), requests temporary credentials via STS `GetSessionToken`.
5. Sends the temporary credentials to the [AWS federation endpoint](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html) to obtain a sign-in token.
6. Constructs a pre-authenticated console URL and opens it in your browser.

## Install

```bash
go install github.com/eculver/go-play/cmd/aws-console@latest
```

## Usage

```
aws-console [flags]

Flags:
  -p, --profile string   AWS profile to use (defaults to AWS_PROFILE env var)
  -h, --help             help for aws-console
```

### Examples

```bash
# Use the profile from AWS_PROFILE env var
export AWS_PROFILE=my-profile
aws-console

# Specify a profile explicitly
aws-console -p my-profile
```

## Prerequisites

- Go 1.21+ (to build)
- The [AWS CLI](https://aws.amazon.com/cli/) must be installed and on your `PATH` (used for SSO login fallback)
- A configured AWS profile in `~/.aws/config` (SSO, IAM user, or assume-role)
