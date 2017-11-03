package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	duration := "4320h"                                                     // 6 months = 6 * 30 * 24
	expiredTime := time.Date(2017, time.January, 10, 12, 0, 0, 2, time.UTC) // expired
	notExpiredTime := time.Now().UTC()                                      // not expired

	delta, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalf("Could not parse duration (%s): %v", duration, err)
	}

	expiresOn := time.Now().Add(delta)
	fmt.Printf("expiration: %s\n", expiresOn)
	fmt.Printf("time until expiration: %s\n", time.Until(expiresOn))

	fmt.Printf("is %s more than %s old?: %t\n", expiredTime, duration, time.Since(expiredTime) > delta)
	fmt.Printf("is %s more than %s old?: %t\n", notExpiredTime, duration, time.Since(notExpiredTime) > delta)
}
