package problem162

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindPeakElement(t *testing.T) {
	testFindPeakElement(t, findPeakElement)
}

func testFindPeakElement(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  map[int]struct{}
	}{
		// preset cases
		{[]int{1, 2, 3, 1}, map[int]struct{}{2: {}}},
		{[]int{1, 2, 1, 3, 5, 6, 4}, map[int]struct{}{1: {}, 5: {}}},
		// common cases
		{[]int{7, 6, 4, 5, 6, 7, 8}, map[int]struct{}{0: {}, 6: {}}},
		{[]int{-7, -6, -8, 5, 8, 6, 3, 9, 8}, map[int]struct{}{1: {}, 4: {}, 7: {}}},
		{[]int{7, 8, 4}, map[int]struct{}{1: {}}},
		{[]int{7, 6, 4, 3, 2, -1, -5}, map[int]struct{}{0: {}}},
		// corner cases
		{[]int{7, 6, 4}, map[int]struct{}{0: {}}},
		{[]int{-1}, map[int]struct{}{0: {}}},
		{[]int{-1, 0}, map[int]struct{}{1: {}}},
		{[]int{0, -1}, map[int]struct{}{0: {}}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if _, ok := testCase.ans[ans]; !ok {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
