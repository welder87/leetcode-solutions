package problem1929

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int) []int

func TestGetConcatenation(t *testing.T) {
	testGetConcatenation(t, getConcatenation)
}

func TestGetConcatenationV1(t *testing.T) {
	testGetConcatenation(t, getConcatenationV1)
}

func testGetConcatenation(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		// preset cases
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
		// common cases
		{[]int{1, 10, 5, 3}, []int{1, 10, 5, 3, 1, 10, 5, 3}},
		// corner cases
		{[]int{1}, []int{1, 1}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
