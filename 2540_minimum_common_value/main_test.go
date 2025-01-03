package problem2540

import (
	"fmt"
	"testing"
)

func TestGetCommon(t *testing.T) {
	testGetCommon(t, getCommon)
}

func testGetCommon(t *testing.T, function fn) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		ans   int
	}{
		{[]int{1, 2, 3}, []int{2, 4}, 2},
		{[]int{1, 2, 3, 6}, []int{2, 3, 4, 5}, 2},
		{[]int{3, 6}, []int{4, 5}, -1},
		{[]int{1, 2, 3, 6}, []int{5, 6, 7, 10, 15}, 6},
		{[]int{2, 2, 2, 2}, []int{1, 1, 1, 2}, 2},
		{[]int{2, 2, 2, 2}, []int{1, 1, 1, 1}, -1},
		{[]int{3}, []int{5}, -1},
		{[]int{3}, []int{3}, 3},
		{[]int{1, 2, 3, 6}, []int{4, 5}, -1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums1, testCase.nums2)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums1, testCase.nums2)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int, []int) int
