package problem167

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int, int) []int

func TestTwoSum(t *testing.T) {
	testTwoSum(t, twoSum)
}

func TestTwoSumV1(t *testing.T) {
	testTwoSum(t, twoSumV1)
}

func TestTwoSumV2(t *testing.T) {
	testTwoSum(t, twoSumV2)
}

func testTwoSum(t *testing.T, function fn) {
	testCases := []struct {
		nums   []int
		target int
		ans    []int
	}{
		// preset cases
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		// common cases
		{[]int{-15, -3, 0, 11, 15}, 8, []int{2, 4}},
		{[]int{-15, -3, 0, 11, 15}, 0, []int{1, 5}},
		{[]int{-15, -3, 0, 11, 15}, 11, []int{3, 4}},
		// corner cases
		{[]int{0, 0}, 0, []int{1, 2}},
		{[]int{-1, -1}, -2, []int{1, 2}},
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
