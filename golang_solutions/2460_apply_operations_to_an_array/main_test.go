package problem2460

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int) []int

func TestApplyOperations(t *testing.T) {
	testApplyOperations(t, applyOperations)
}

func testApplyOperations(t *testing.T, fucntion fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		// preset cases
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{
			[]int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272},
			[]int{1694, 399, 832, 1758, 412, 206, 272, 0, 0, 0, 0, 0, 0, 0},
		},
		// common cases
		{[]int{2, 2, 2, 2, 3, 1}, []int{4, 4, 3, 1, 0, 0}},
		{[]int{1, 3, 1, 3, 4, 0}, []int{1, 3, 1, 3, 4, 0}},
		{[]int{0, 3, 1, 3, 4, 2}, []int{3, 1, 3, 4, 2, 0}},
		// corner cases
		{[]int{0, 0}, []int{0, 0}},
		{[]int{2, 2}, []int{4, 0}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := fucntion(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
