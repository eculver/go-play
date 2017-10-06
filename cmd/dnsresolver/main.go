package main

import (
	"fmt"
	"log"
	"net"
)

// to run:
// GODEBUG=netdns=go go run cmd/dnsresolver/main.go
func main() {
	url := "www.uber.com."
	// url := "yahoo.com"
	addrs, err := net.LookupHost(url)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	fmt.Println(addrs)
}
