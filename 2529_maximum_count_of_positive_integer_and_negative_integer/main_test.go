package problem2529

import (
	"fmt"
	"testing"
)

func TestMaximumCount(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{-2, -1, -1, 1, 2, 3},
			ans:  3,
		},
		{
			nums: []int{-3, -2, -1, 0, 0, 1, 2},
			ans:  3,
		},
		{
			nums: []int{5, 20, 66, 1314},
			ans:  4,
		},
		{
			nums: []int{1},
			ans:  1,
		},
		{
			nums: []int{-1},
			ans:  1,
		},
		{
			nums: []int{-2, -1, -1},
			ans:  3,
		},
		{
			nums: []int{-2, -2, -2, -2, -1, 1, 2, 3, 5},
			ans:  5,
		},
		{
			nums: []int{-2, 0, 0},
			ans:  1,
		},
		{
			nums: []int{0},
			ans:  0,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0, 0},
			ans:  0,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0, 1},
			ans:  1,
		},
		{
			nums: []int{-1, 0, 0, 0, 0, 0, 1},
			ans:  1,
		},
		{
			nums: []int{-2, -2, -1, -1, -1, -1, 5},
			ans:  6,
		},
		{
			nums: []int{-2, -2, -1, -1, -1, 3, 5},
			ans:  5,
		},
		{
			nums: []int{-2, -2, -1, -1, -1, 0, 5},
			ans:  5,
		},
		{
			nums: []int{-2, -2, -1, -1, -1, 0},
			ans:  5,
		},
		{
			nums: []int{0, 1, 2, 5},
			ans:  3,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := maximumCount(testCase.nums)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
