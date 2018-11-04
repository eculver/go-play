package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

type Decoder interface {
	Decode(interface{}) error
}

type Encoder interface {
	Encode(interface{}) error
}

type Base64Decoder struct {
	src io.Reader
	dst io.Writer
}

func (d *Base64Decoder) Decode(v interface{}) error {
	// bs := v.([]byte)
	return nil
}

func NewBase64Decoder(src io.Reader, dst io.Writer) Decoder {
	return &Base64Decoder{
		src: base64.NewDecoder(base64.StdEncoding, src),
		dst: dst,
	}
}

type StringEncoder struct {
	buf *bytes.Buffer
	out io.Writer
}

func NewStringEncoder(out io.Writer) *StringEncoder {
	return &StringEncoder{
		buf: new(bytes.Buffer),
		out: out,
	}
}

func (e *StringEncoder) Write(v []byte) (int, error) {
	return e.buf.WriteString(string(v))
}

func (e *StringEncoder) Close() error {
	_, err := e.buf.WriteTo(e.out)
	return err
}

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

	input := new(bytes.Buffer)
	for _, val := range args {
		input.WriteString(val + "\n")
	}

	reader, writer := io.Pipe()
	var decoder io.Reader
	var encoder io.WriteCloser

	switch *from {
	case "base64":
		decoder = base64.NewDecoder(base64.StdEncoding, input)
	case "ascii":
		decoder = input
	default:
		fmt.Println("ERROR: unknown input format")
		fmt.Println(usage())
		os.Exit(1)
	}

	switch *to {
	case "base64":
		encoder = base64.NewEncoder(base64.StdEncoding, writer)
	case "ascii":
		encoder = NewStringEncoder(os.Stdout)
	default:
		fmt.Println("ERROR: unknown output format")
		fmt.Println(usage())
		os.Exit(1)
	}

	go func() {
		io.Copy(encoder, decoder)
		encoder.Close()
		pwriter.Close()
	}()

	io.Copy(os.Stdout, preader)
	preader.Close()
}
