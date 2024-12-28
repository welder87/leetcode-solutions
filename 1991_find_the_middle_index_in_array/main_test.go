package problem1991

import (
	"fmt"
	"testing"
)

func TestFindMiddleIndex(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{[]int{2, 3, -1, 8, 4}, 3},
		{[]int{1, -1, 4}, 2},
		{[]int{2, 5}, -1},
		{[]int{5}, 0},
		{[]int{2, 3, -6, 8}, 1},
		{[]int{4, 3, -6, 8, 2}, 1},
		{[]int{2, 6, -6, 8, 2}, 3},
		{[]int{15, -9, 3, 5, 1}, 0},
		{[]int{-9, 3, 5, 1, 15}, 4},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := findMiddleIndex(testCase.nums)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
