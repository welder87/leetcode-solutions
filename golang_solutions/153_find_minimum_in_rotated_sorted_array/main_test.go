package problem153

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindMin(t *testing.T) {
	testFindMin(t, findMin)
}

func TestFindMinV1(t *testing.T) {
	testFindMin(t, findMinV1)
}

func testFindMin(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		// common cases
		{[]int{4, 6, 7, 8, 9, 0, 1, 2}, 0},
		{[]int{17, 11, 13, 15}, 11},
		{[]int{2, 3, 4, 5, 1}, 1},
		{[]int{4, 5, 1, 2, 3}, 1},
		{[]int{5, 1, 2, 3, 4}, 1},
		{[]int{5, 8, -7, -3, 0}, -7},
		// corner cases
		{[]int{1, 0}, 0},
		{[]int{3}, 3},
		{[]int{-3}, -3},
		{[]int{0}, 0},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
