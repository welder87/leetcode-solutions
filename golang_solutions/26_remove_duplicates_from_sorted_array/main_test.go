package problem26

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int) int

func TestRemoveDuplicates(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicates)
}

func TestRemoveDuplicatesV1(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicatesV1)
}

func testRemoveDuplicates(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		arr  []int
		ans  int
	}{
		// preset cases
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		// common cases
		{[]int{-5, -1, -1, 0, 0, 1, 2, 2, 4, 4}, []int{-5, -1, 0, 1, 2, 4}, 6},
		{[]int{1, 2, 2}, []int{1, 2}, 2},
		{[]int{1, 1, 1, 1, 2}, []int{1, 2}, 2},
		// corner cases
		{[]int{-1, -1, -1}, []int{-1}, 1},
		{[]int{0}, []int{0}, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
			if !reflect.DeepEqual(testCase.nums[:ans], testCase.arr) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.arr)
			}
		})
	}
}
