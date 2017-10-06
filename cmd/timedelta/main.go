package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	delta, err := time.ParseDuration("4320h") // 6 months = 6 * 30 * 24
	if err != nil {
		log.Fatalf("Could not parse duration: %v", err)
	}
	// delta := time.Duration(4320) * time.Hour
	expiresOn := time.Now().Add(delta)
	fmt.Printf("expiration: %s\n", expiresOn)
	fmt.Printf("time until expiration: %s\n", time.Until(expiresOn))
}
