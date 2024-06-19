package maxksumpairs_test

import (
	"fmt"
	"testing"

	"github.com/eculver/go-play/pkg/challenge/maxksumpairs"
)

func Test_MaxKSumPairs(t *testing.T) {
	tcs := []struct {
		in  []int
		k   int
		out int
	}{
		{
			in:  []int{1, 2, 3, 4},
			k:   5,
			out: 2,
		},
		{
			in:  []int{3, 1, 3, 4, 3},
			k:   6,
			out: 1,
		},
	}

	// test the brute force solution
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Brute solution test case %d", i), func(t *testing.T) {
			out := maxksumpairs.Brute(tc.in, tc.k)
			if out != tc.out {
				t.Errorf("Expected %d, got %d", tc.out, out)
			}
		})
	}
	// test the map-based solution
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Map-based solution test case %d", i), func(t *testing.T) {
			out := maxksumpairs.Map(tc.in, tc.k)
			if out != tc.out {
				t.Errorf("Expected %d, got %d", tc.out, out)
			}
		})
	}
}
