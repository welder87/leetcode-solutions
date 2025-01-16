package problem153

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 4, 5, 1, 2},
			ans:  1,
		},
		{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			ans:  0,
		},
		{
			nums: []int{11, 13, 15, 17},
			ans:  11,
		},
		{
			nums: []int{4, 6, 7, 8, 9, 0, 1, 2},
			ans:  0,
		},
		{
			nums: []int{17, 11, 13, 15},
			ans:  11,
		},
		{
			nums: []int{1},
			ans:  1,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := findMin(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
