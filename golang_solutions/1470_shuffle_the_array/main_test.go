package problem1470

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int, int) []int


func TestShuffle(t *testing.T,) {
	testShuffle(t, shuffle)
}

func testShuffle(t *testing.T, function fn) {
	testCases := []struct {
		nums   []int
		n int
		ans []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
		{[]int{1, 2}, 1, []int{1, 2}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.n)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums, testCase.n)
			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
