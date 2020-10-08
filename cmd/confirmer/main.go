// package main is a simple demonstration of how to use the confirmer library
package main

import (
	"fmt"
	"os"

	"github.com/eculver/go-play/pkg/confirmer"
)

func main() {
	if !confirmer.Confirm("Continue?", os.Stdin) {
		// not confirmed
		fmt.Println("not confirmed!")
		os.Exit(1)
	}
	// confirmed
	fmt.Println("confirmed!")
}
