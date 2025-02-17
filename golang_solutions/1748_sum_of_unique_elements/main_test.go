package problem1748

import (
	"fmt"
	"testing"
)

func TestSumOfUnique(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{1, 2, 3, 2}, 4},
		{[]int{1, 1, 1, 1, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 15},
		// common cases
		{[]int{1, 1, 1, 5}, 5},
		{[]int{1, 5, 5, 5}, 1},
		// corner cases
		{[]int{1}, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := sumOfUnique(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
