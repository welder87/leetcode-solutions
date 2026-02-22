package problem1508

import "sort"

// Time complexity: O(n*n + m log m + k). Space complexity: O(~n*n + m).
func rangeSum(nums []int, n int, left int, right int) int {
	ar := make([]int, 0, n*(n+1)/2)
	sm := 0
	for i := range nums {
		sm := 0
		for j := i; j < len(nums); j++ {
			sm += nums[j]
			ar = append(ar, sm)
		}
	}
	sort.Ints(ar)
	sm = 0
	mod := 10*10*10*10*10*10*10*10*10 + 7
	for i := left - 1; i < right; i++ {
		sm = (sm + ar[i]) % mod
	}
	return sm
}
