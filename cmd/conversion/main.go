package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() string {
	return fmt.Sprintf("usage: ./%s -from string -to bytes 'abcdedf'", os.Args[0])
}

func main() {
	from := flag.String("from", "string", "Input format. Available input formats are ...")
	to := flag.String("to", "base64", "Output format. Available output formats are ...")
	flag.Parse()
	args := flag.Args()

	if from == nil {
		fmt.Println("ERROR: -from flag is required")
		fmt.Println(usage())
		os.Exit(1)
	}
	if to == nil {
		fmt.Println("ERROR: -to flag is required")
		fmt.Println(usage())
		os.Exit(2)
	}
	if len(args) == 0 {
		fmt.Println("ERROR: must pass something to convert")
		fmt.Println(usage())
		os.Exit(1)
	}

	var (
		decoder io.Reader
		encoder io.Writer
	)

	reader, writer := io.Pipe()

	switch *from {
	case "base64":
		decoder = base64.NewDecoder(base64.StdEncoding, reader)
	case "ascii":
		decoder = bytes.NewBuffer([]byte{})
	default:
		fmt.Println("ERROR: unknown input format")
		fmt.Println(usage())
		os.Exit(1)
	}

	switch *to {
	case "base64":
		encoder = base64.NewEncoder(base64.StdEncoding, writer)
	case "ascii":
		encoder = bytes.NewBuffer([]byte{})
	default:
		fmt.Println("ERROR: unknown output format")
		fmt.Println(usage())
		os.Exit(1)
	}

	for _, val := range args {
		fmt.Fprintf(writer, val)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	fmt.Print(buf.String())
}
