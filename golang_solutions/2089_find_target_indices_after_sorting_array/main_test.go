package problem2089

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int, int) []int

func TestTargetIndices(t *testing.T) {
	testTargetIndices(t, targetIndices)
}

func TestTargetIndicesV1(t *testing.T) {
	testTargetIndices(t, targetIndicesV1)
}

func testTargetIndices(t *testing.T, function fn) {
	testCases := []struct {
		nums   []int
		target int
		ans    []int
	}{
		// preset cases
		{[]int{1, 2, 5, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 5, 2, 3}, 3, []int{3}},
		{[]int{1, 2, 5, 2, 3}, 5, []int{4}},
		// common cases
		{[]int{1, 2, 5, 2, 3, 2}, 2, []int{1, 2, 3}},
		{[]int{1, 2, 5, 5, 3, 5}, 5, []int{3, 4, 5}},
		{[]int{1, 2, 1, 1, 3, 5}, 1, []int{0, 1, 2}},
		{[]int{1, 2, 1, 1, 3, 5}, 4, []int{}},
		// corner cases
		{[]int{2, 2, 2}, 2, []int{0, 1, 2}},
		{[]int{2, 2, 2}, 3, []int{}},
		{[]int{2}, 2, []int{0}},
		{[]int{1}, 2, []int{}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums, testCase.target)
			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
