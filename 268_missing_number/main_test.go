package problem268

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 0, 1},
			ans:  2,
		},
		{
			nums: []int{0, 1},
			ans:  2,
		},
		{
			nums: []int{1},
			ans:  0,
		},
		{
			nums: []int{0},
			ans:  1,
		},
		{
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			ans:  8,
		},
		{
			nums: []int{9, 6, 4, 2, 3, 5, 7, 8, 1},
			ans:  0,
		},
		{
			nums: []int{9, 6, 4, 2, 3, 5, 7, 8, 0, 1},
			ans:  10,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := missingNumber(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
