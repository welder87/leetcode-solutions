package problem1636

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int) []int

func TestFrequencySort(t *testing.T) {
	testFrequencySort(t, frequencySort)
}

func testFrequencySort(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{[]int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{2, 3, 1, 3, 2}, []int{3, 1, 1, 2, 2, 2}},
		{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
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
