package problem349

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	testIntersection(t, intersection)
}

func TestIntersectionV1(t *testing.T) {
	testIntersection(t, intersectionV1)
}

func testIntersection(t *testing.T, function fn) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		ans   [][]int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, [][]int{{2}}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, [][]int{{9, 4}, {4, 9}}},
		{[]int{1, 2, 3}, []int{4, 6, 5}, [][]int{{}}},
		{[]int{}, []int{}, [][]int{{}}},
		{[]int{1, 5, 1, 3, 5, 7}, []int{1, 3}, [][]int{{1, 3}, {3, 1}}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums1, testCase.nums2)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums1, testCase.nums2)
			isFound := false
			for _, expectedAns := range testCase.ans {
				if reflect.DeepEqual(expectedAns, ans) {
					isFound = true
				}
			}
			if !isFound {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int, []int) []int
