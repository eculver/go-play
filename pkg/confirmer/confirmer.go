package confirmer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Confirmer confirms user input
type Confirmer interface {
	Confirm(msg string, reader io.Reader) bool
}

var DefaultAccept = []string{"Y"}
var DefaultDeny = []string{"n"}
var DefaultConfirmer = New(DefaultAccept, DefaultDeny, 3)

func Confirm(msg string, reader io.Reader) bool {
	return DefaultConfirmer.Confirm(msg, reader)
}

// New returns a new Confirmer for reading acceptable or deniable
// input for the given number of tries before giving up.
func New(accept []string, deny []string, tries int) Confirmer {
	return &confirmer{
		accept: accept,
		deny:   deny,
		tries:  tries,
	}
}

type confirmer struct {
	accept []string
	deny   []string
	tries  int
}

// confirm displays a prompt `s` to the user and returns a bool indicating yes / no
// If the lowercased, trimmed input begins with anything other than 'y', it returns false
// It accepts an int `tries` representing the number of attempts before returning false
func (c confirmer) Confirm(msg string, reader io.Reader) bool {
	opts := append(c.accept, c.deny...)
	optsStr := strings.Join(opts, "/")
	r := bufio.NewReader(reader)

	for ; c.tries > 0; c.tries-- {
		fmt.Printf("%s [%s]: ", msg, optsStr)

		res, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("ERROR: could not read input: %s\n", err)
			return false
		}
		// empty input (i.e. "\n")
		if len(res) < 2 {
			fmt.Println("")
			continue
		}
		res = strings.TrimSpace(res)

		for _, a := range c.accept {
			if res == a {
				// fmt.Println("")
				return true
			}
		}
		for _, d := range c.deny {
			if res == d {
				// fmt.Println("")
				return false
			}
		}
		fmt.Println("")
	}
	return false
}
