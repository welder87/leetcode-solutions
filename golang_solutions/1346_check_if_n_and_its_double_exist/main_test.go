package problem1346

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int) bool

func TestCheckIfExist(t *testing.T) {
	testTestCheckIfExist(t, checkIfExist)
}

func testTestCheckIfExist(t *testing.T, function fn) {
	testCases := []struct {
		arr []int
		ans bool
	}{
		// preset cases
		{[]int{10, 2, 5, 3}, true},
		{[]int{3, 1, 7, 11}, false},
		// common cases
		{[]int{1, -5, -10, 12, 0, 6}, true},
		{[]int{-10, -1, 0, 15, 1, 2}, true},
		{[]int{-10, -1, 0, 15, 1, 3}, false},
		{[]int{-10, -1, 0, 0, 15, 1, 3}, true},
		// corner cases
		{[]int{-4, -2}, true},
		{[]int{-4, 2}, false},
		{[]int{2, 4}, true},
		{[]int{-2, 4}, false},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.arr)
		t.Run(testName, func(t *testing.T) {
			result := function(testCase.arr)
			if !reflect.DeepEqual(result, testCase.ans) {
				t.Errorf("got %v, want %v", result, testCase.ans)
			}
		})
	}
}
