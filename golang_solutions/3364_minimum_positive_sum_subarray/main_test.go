package problem3364

import (
	"fmt"
	"testing"
)

func TestMinimumSumSubarray(t *testing.T) {
	testCases := []struct {
		nums []int
		l    int
		r    int
		ans  int
	}{
		// preset cases
		{[]int{3, -2, 1, 4}, 2, 3, 1},
		{[]int{-2, 2, -3, 1}, 2, 3, -1},
		{[]int{1, 2, 3, 4}, 2, 4, 3},
		// common cases
		{[]int{3, -2, 1, 4}, 1, 1, 1},
		{[]int{3, -3, 2, 4}, 1, 3, 2},
		// corner cases
		{[]int{1}, 0, 0, -1},
		{[]int{0}, 0, 0, -1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v %v", testCase.nums, testCase.l, testCase.r)
		t.Run(testName, func(t *testing.T) {
			ans := minimumSumSubarray(testCase.nums, testCase.l, testCase.r)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
