package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const baseURL = "https://crawl.codeferret.net"

var (
	linkRegex = regexp.MustCompile(`href\s*=\s*[\'"]?([^\'" >]+)`)
	urlRegex  = regexp.MustCompile(`(.*)(\/\w+\.html)`)
)

func main() {
	start := fmt.Sprintf("%s/1.html", baseURL)
	urls := getLinks(start)
	followed := []string{start}
	for _, url := range urls {
		if shouldFollow(url, followed) {
			followed = append(followed, url)
			for _, u := range follow(url, followed) {
				urls = append(urls, u)
			}
		}
	}
	fmt.Printf("links: %v\n", urls)
}

func follow(url string, followed []string) []string {
	urls := getLinks(url)
	followed = append(followed, url)
	for _, url := range urls {
		if shouldFollow(url, followed) {
			for _, u := range follow(url, followed) {
				urls = append(urls, u)
			}
		}
	}
	return urls
}

func getLinks(url string) []string {
	fmt.Println("following url " + url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP Error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	matches := linkRegex.FindAllString(string(body), -1)
	urls := make([]string, len(matches))
	for i, m := range matches {
		urls[i] = fmt.Sprintf("%s%s", baseURL, urlRegex.ReplaceAllString(m, "$2"))
	}
	return urls
}

func shouldFollow(url string, followed []string) bool {
	for _, u := range followed {
		if u == url {
			return false
		}
	}
	return true
}
