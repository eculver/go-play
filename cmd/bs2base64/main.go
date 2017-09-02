package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var argsRegex = regexp.MustCompile(`\d+`)

func usage() string {
	return fmt.Sprintf("usage: ./%s \"63 109 71 252 183 247 73 237 176 74 104 198 29 250 5 113\"", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("invalid arguments\n%s\n", usage())
	}

	matches := argsRegex.FindAllString(os.Args[1], -1)

	if len(matches) < 16 {
		log.Fatalf("argument must contain 16 bytes: %v, args: %v", matches, os.Args[1])
	}

	bs := make([]byte, 16)
	for idx, b := range matches {
		b, err := strconv.ParseUint(b, 10, 64)
		if err != nil {
			log.Fatalf("could not parse uint: %v", b)
		}
		bs[idx] = byte(uint8(b))
	}
	fmt.Println(base64.URLEncoding.EncodeToString(bs))
}
