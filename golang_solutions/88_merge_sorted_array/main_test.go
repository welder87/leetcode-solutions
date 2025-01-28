package problem88

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int, int, []int, int)

func TestMerge(t *testing.T) {
	testMerge(t, merge)
}

func TestMergeV1(t *testing.T) {
	testMerge(t, mergeV1)
}

func testMerge(t *testing.T, function fn) {
	testCases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		ans   []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 3, 0, 0, 0}, 4, []int{1, 2, 3}, 3, []int{1, 1, 2, 2, 3, 3, 3}},
		{[]int{4, 5, 5, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 5}},
		{[]int{1, 2, 4, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 4, 4, 5, 6}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 2, 0, 0, 0}, 3, []int{1, 5, 6}, 3, []int{1, 1, 2, 2, 5, 6}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v %v %v", testCase.nums1, testCase.m, testCase.nums2, testCase.n)
		t.Run(testName, func(t *testing.T) {
			function(testCase.nums1, testCase.m, testCase.nums2, testCase.n)
			if !reflect.DeepEqual(testCase.ans, testCase.nums1) {
				t.Errorf("got %v, want %v", testCase.nums1, testCase.ans)
			}
		})
	}
}
