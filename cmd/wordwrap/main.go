package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func main() {
	tcs := []struct {
		input    string
		llength  int
		expected []string
	}{
		{"a b c d", 1, []string{"a", "b", "c", "d"}},
	}

	for _, tc := range tcs {
		res, err := wordwrap(tc.input, tc.llength)
		if tc.expected == nil || len(tc.expected) == 0 && err == nil {
			log.Fatalf("FAIL - for input (%s, %d), expected error", tc.input, tc.llength)
		}
		if len(res) != len(tc.expected) {
			log.Fatalf("FAIL - for input (%s, %d), got: %d, expected: %d", tc.input, tc.llength, len(res), len(tc.expected))
		}
	}
	fmt.Println("SUCCESS")
}

func wordwrap(input string, llength int) ([]string, error) {
	tokens := strings.Split(input, " ")
	ret := []string{}

	if len(tokens) == 0 {
		return nil, errors.New("empty tokens")
	}

	curr := tokens[0]
	line := ""

	for i := 0; i < len(tokens)-1; i++ {
		if len(curr) > llength {
			return nil, fmt.Errorf("overflow on token %s", curr)
		}
		maybe := fmt.Sprintf("%s %s", curr, tokens[i+1])
		if len(maybe) < llength {
			line = maybe
			curr = tokens[i+1]
			continue
		}
		ret = append(ret, line)
	}
	return ret, nil
}
