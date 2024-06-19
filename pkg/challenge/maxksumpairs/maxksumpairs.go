// package maxksumpairs provides the methods for solving the Max K Sum Pairs challenge.
// The challenge is to find the maximum number of pairs of integers in a list that sum to k.
// The original input should be modified to remove the pairs that sum to k and the number of operations (aka pairs found) should be returned
//
// Example 1:
// Input: nums = [1, 2, 3, 4], k = 5
// Output: 2
// Explanation: The optimal solution is to remove [1, 4] and then [2, 3] and then return 2 for the number of operations.
//
// Example 2:
// Input: nums = [3, 1, 3, 4, 3], k = 6
// Output: 1
// Explanation: Remove [3, 3], leaving [1, 4, 3] for which there are no more pairs that sum to 6.
package maxksumpairs

import "sort"

// Brute solves the problem by rolling the array and comparing each pair.
func Brute(nums []int, k int) int {
	// slice has to be sorted other
	sort.Ints(nums)

	// go through from the front to the back finding pairs that sum to k
	count := 0
	i := 0
	j := len(nums) - 1
	for {
		if j <= i {
			break
		}
		sum := nums[i] + nums[j]
		if sum == k {
			count++
			i++
			j--
			continue
		}
		if sum > k {
			j--
			continue
		}
		i++
	}

	return count
}

// Map uses a map to keep track of the pairs and their differences
func Map(nums []int, k int) int {
	count := 0
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	for v, _ := range m {
		seek := k - v
		if _, ok := m[seek]; ok {
			count++
			if m[v] == 1 {
				delete(m, v)
			} else {
				m[v]--
			}

			if m[seek] == 1 {
				delete(m, seek)
			} else {
				m[seek]--
			}
		}
	}
	return count
}
