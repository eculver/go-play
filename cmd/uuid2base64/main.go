package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	uuid "github.com/satori/go.uuid"
)

func usage() string {
	return fmt.Sprintf("usage: ./%s \"ad17f886-1f3b-4e38-a490-c60421bf0455\"", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("invalid arguments\n%s\n", usage())
	}
	u, err := uuid.FromString(os.Args[1])
	if err != nil {
		log.Fatalf("argument is invalid: %v, args: %v", err, os.Args[1])
	}
	fmt.Println(base64.URLEncoding.EncodeToString(u.Bytes()))
}
