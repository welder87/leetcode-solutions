package problem485

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testFindMaxConsecutiveOnes(t, findMaxConsecutiveOnes)
}

func testFindMaxConsecutiveOnes(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			ans:  3,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1},
			ans:  2,
		},
		{
			nums: []int{1, 1, 1, 1, 1, 1},
			ans:  6,
		},
		{
			nums: []int{0, 0, 0, 0, 0, 0},
			ans:  0,
		},
		{
			nums: []int{0},
			ans:  0,
		},
		{
			nums: []int{1},
			ans:  1,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1, 1},
			ans:  2,
		},
		{
			nums: []int{1, 1, 1, 0, 1, 1, 0, 1, 1},
			ans:  3,
		},
		{
			nums: []int{1, 1, 0, 1, 0, 1, 1, 1, 0},
			ans:  3,
		},
		{
			nums: []int{0, 1, 0, 0, 1, 1, 0, 1, 0, 0},
			ans:  2,
		},
		{
			nums: []int{0, 1, 0, 1, 0, 1, 0},
			ans:  1,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1, 1, 1},
			ans:  3,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) int
