package problem2215

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDifference(t *testing.T) {
	testFindDifference(t, findDifference)
}

func testFindDifference(t *testing.T, fucntion fn) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		ans  [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums1, testCase.nums2)
		t.Run(testName, func(t *testing.T) {
			ans := fucntion(testCase.nums1, testCase.nums2)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int, []int) [][]int
