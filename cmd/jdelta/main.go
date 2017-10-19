package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	tcs := []struct {
		in       []int
		delta    int
		expected int
	}{
		{[]int{}, 0, 0},         // empty
		{[]int{1}, 0, 0},        // size: 1
		{[]int{1, 2, 10}, 0, 0}, // size: n

		{[]int{1, 2, 3, 4}, 0, 0}, // sorted ASC, zero delta, no matches
		{[]int{4, 3, 2, 1}, 0, 0}, // sorted DESC, zero delta, no matches
		{[]int{2, 1, 3, 4}, 0, 0}, // unsorted, zero delta, no matches

		{[]int{1, 2, 3, 4}, 10, 0}, // sorted ASC, non-zero delta, no matches
		{[]int{4, 3, 2, 1}, 10, 0}, // sorted DESC, non-zero delta, no matches
		{[]int{2, 1, 3, 4}, 10, 0}, // unsorted, non-zero delta, no matches

		{[]int{1, 2, 3, 4}, 3, 1}, // sorted ASC, non-zero delta, single match
		{[]int{4, 3, 2, 1}, 3, 1}, // sorted DESC, non-zero delta, single match
		{[]int{2, 1, 3, 4}, 3, 1}, // unsorted, non-zero delta, single match

		{[]int{1, 2, 3, 4}, 2, 2}, // sorted ASC, non-zero delta, n-matches
		{[]int{4, 3, 2, 1}, 2, 2}, // sorted DESC, non-zero delta, n-matches
		{[]int{2, 1, 3, 4}, 2, 2}, // unsorted, non-zero delta, n-matches

		{[]int{1, 12, 5, 3, 4, 2}, 2, 3}, // from example
	}

	for _, tc := range tcs {
		res := jdelta(tc.in, tc.delta)
		if res != tc.expected {
			log.Fatalf("FAILED FOR %v with delta %d - got: %d, expected: %d", tc.in, tc.delta, res, tc.expected)
		}
	}
	fmt.Printf("SUCCESS - YAAAAY!")
}

// technically, this is O(n + n) ~ O(2n) if we can't pre-compute map
func jdelta(in []int, delta int) int {
	if delta < 1 {
		return 0 // delta must be positive, non-zero
	}
	matches := make(map[string]bool) // map of all _unique_ matches
	mlow := make(map[string]int)     // map of "low" matches (i - delta)
	mhigh := make(map[string]int)    // map of "high" matches (i + delta)

	// build a map of *possible* matches per index value, hopefully pre-computed?
	for i, v := range in {
		mlow[strconv.Itoa(v-delta)] = i
		mhigh[strconv.Itoa(v+delta)] = i
	}

	// count only if not already matched
	// matches are unique pairs of (low, high) in that order -- MATH
	for i, v := range in {
		// does v exist in low map?
		if idx, ok := mlow[strconv.Itoa(v)]; ok {
			matches[getMatchIdx(i, idx)] = true
			continue
		}
		// does v exist in high map?
		if idx, ok := mhigh[strconv.Itoa(v)]; ok {
			matches[getMatchIdx(i, idx)] = true
		}
	}
	return len(matches)
}

// for deduping matches
func getMatchIdx(i, idx int) string {
	matchidx := fmt.Sprintf("(%d,%d)", i, idx)
	if i > idx {
		matchidx = fmt.Sprintf("(%d,%d)", idx, i)
	}
	return matchidx
}
