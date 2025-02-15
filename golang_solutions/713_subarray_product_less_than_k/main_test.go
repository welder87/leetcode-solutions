package problem713

import (
	"fmt"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	testCases := []struct {
		nums   []int
		k int
		ans int
	}{
		// preset cases
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		// common cases
		{[]int{10, 5, 2, 6}, 13, 6},
		{[]int{10, 5, 2, 12, 6}, 12, 5},
		{[]int{10, 5, 2, 6}, 1000, 10},
		// corner cases
		{[]int{1, 1}, 1, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, 1},
		{[]int{1}, 0, 0},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := numSubarrayProductLessThanK(testCase.nums, testCase.k)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
