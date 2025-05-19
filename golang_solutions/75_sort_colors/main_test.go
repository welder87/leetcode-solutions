package problem75

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int)

func TestSortColors(t *testing.T) {
	testSortColors(t, sortColors)
}

func TestSortColorsV1(t *testing.T) {
	testSortColors(t, sortColorsV1)
}

func testSortColors(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		// preset cases
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		// common cases
		{[]int{2, 1, 1, 1, 0}, []int{0, 1, 1, 1, 2}},
		// corner cases
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1}, []int{1}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			function(testCase.nums)
			if !reflect.DeepEqual(testCase.nums, testCase.ans) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.ans)
			}
		})
	}
}
