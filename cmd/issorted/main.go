package main

import (
	"fmt"
	"log"
	"math"
)

func isSorted(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] <= s[1]
	}
	pivot := int(math.Floor(float64(len(s) / 2)))
	fmt.Printf("pivot: %d\n", pivot)
	fmt.Printf("part1: %v, part2: %v\n", s[:pivot], s[pivot:])
	return isSorted(s[:pivot]) && isSorted(s[pivot:])
}

func main() {
	tcs := []struct {
		in       string
		expected bool
	}{
		{
			in:       "abcdefg",
			expected: true,
		},
		{
			in:       "gfedc",
			expected: false,
		},
		{
			in:       "a",
			expected: true,
		},
		{
			in:       "aaaaaaaaaaaaaaa",
			expected: true,
		},
		{
			in:       "zzzzzzzzzzzzzzz",
			expected: true,
		},
		{
			in:       "zzzzzzzzzzzzzza",
			expected: false,
		},
	}
	for _, tc := range tcs {
		if isSorted(tc.in) != tc.expected {
			log.Fatalf("FAILURE: %v != %v", tc.in, tc.expected)
		}
	}
}
