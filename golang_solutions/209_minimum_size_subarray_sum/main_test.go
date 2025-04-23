package problem209

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func(int, []int) int

func TestMinSubArrayLen(t *testing.T) {
	testMinSubArrayLen(t, minSubArrayLen)
}

func testMinSubArrayLen(t *testing.T, function fn) {
	testCases := []struct {
		nums   []int
		target int
		ans    int
	}{
		// preset cases
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 4, 4}, 4, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
		// common cases
		{[]int{2, 1, 3, 1}, 7, 4},
		{[]int{2, 1, 3, 1}, 8, 0},
		{[]int{2, 1, 3, 2}, 8, 4},
		// corner cases
		{[]int{1}, 11, 0},
		{[]int{1}, 1, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.target, testCase.nums)
			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
