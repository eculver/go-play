package main

import "fmt"

type testcase struct {
	input    []string
	expected int
}

var testcases = []testcase{
	{[]string{"10100", "10111", "11111", "10010"}, 4},
	{[]string{"10111", "10111", "11111", "10010"}, 9},
	{[]string{"01"}, 1},
}

func main() {
	for _, tc := range testcases {
		in := make([][]byte, len(tc.input))
		for i, str := range tc.input {
			in[i] = []byte(str)
		}

		res := maximalSquare(in)
		if res != tc.expected {
			fmt.Printf("FAIL: %v - got %d, expected %d\n", tc.input, res, tc.expected)
			continue
		}
		fmt.Println("PASS")
	}
}

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}
	maxsqlen := 0
	prev := 0
	dp := make([]int, cols+1)

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] != 0 {
				tmp := dp[j]
				if matrix[i-1][j-1] == '1' {
					dp[j] = min(min(dp[j-1], prev), dp[j]) + 1
					if dp[j] > maxsqlen {
						maxsqlen = dp[j]
					}
				} else {
					dp[j] = 0
				}
				prev = tmp
			}
		}
	}
	return maxsqlen * maxsqlen

}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
