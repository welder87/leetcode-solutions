package problem1394

import (
	"fmt"
	"testing"
)

func TestFindLucky(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{2, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{2, 2, 2, 3, 3}, -1},
		// common cases
		{[]int{1, 2, 2, 2, 3, 3, 3}, 3},
		{[]int{1, 4, 4, 4, 4, 3, 3, 3}, 4},
		// corner cases
		{[]int{1}, 1},
		{[]int{2}, -1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := findLucky(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
