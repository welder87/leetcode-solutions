package problem697

import (
	"fmt"
	"testing"
)

func TestFindShortestSubArray(t *testing.T) {
	testCases := []struct {
		nums   []int
		ans int
	}{
		// preset cases
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
		// common cases
		{[]int{2, 1, 2, 3, 1, 4, 2}, 7},
		{[]int{0, 3, 1, 2, 2, 3, 1}, 2},
		{[]int{1, 1, 2, 2, 3, 3}, 2},
		// corner cases
		{[]int{1}, 1},
		{[]int{1, 0}, 1},
		{[]int{1, 1}, 2},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := findShortestSubArray(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
