package problem724

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	testPivotIndex(t, pivotIndex)
}

func TestPivotIndexV1(t *testing.T) {
	testPivotIndex(t, pivotIndexV1)
}

func testPivotIndex(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			ans:  -1,
		},
		{
			nums: []int{1, 7, 3, 6, 5, 6},
			ans:  3,
		},
		{
			nums: []int{1, 2, 3},
			ans:  -1,
		},
		{
			nums: []int{2, 1, -1},
			ans:  0,
		},
		{
			nums: []int{1},
			ans:  0,
		},
		{
			nums: []int{-1, 1, 2},
			ans:  2,
		},
		{
			nums: []int{12, -3, 10, 5, -46, -31},
			ans:  -1,
		},
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

type fn func([]int) int
