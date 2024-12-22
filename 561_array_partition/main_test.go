package problem561

import (
	"fmt"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{1, 4, 3, 2},
			ans:  4,
		},
		{
			nums: []int{6, 2, 6, 5, 1, 2},
			ans:  9,
		},
		{
			nums: []int{6, 9, 6, 3, 8, 4},
			ans:  17,
		},
		{
			nums: []int{-6, 2, 6, 5, 1, 2, 6, -5},
			ans:  3,
		},
		{
			nums: []int{-6, 2, 6, 0, 1, 2, 6, -5},
			ans:  2,
		},
		{
			nums: []int{-6, 10},
			ans:  -6,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := arrayPairSum(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
