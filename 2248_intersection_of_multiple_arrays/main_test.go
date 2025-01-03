package problem2248

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	testIntersection(t, intersection)
}

func testIntersection(t *testing.T, fucntion fn) {
	testCases := []struct {
		nums [][]int
		ans  []int
	}{
		{
			[][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}},
			[]int{3, 4},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{},
		},
		{
			[][]int{{1, 2, 3}, {1, 3, 2}, {3, 1, 2}},
			[]int{1, 2, 3},
		},
		{
			[][]int{{3, 1, 2, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[][]int{{1, 1, 1}, {1, 5, 6}},
			[]int{1},
		},
		{
			[][]int{{44, 21, 48}},
			[]int{21, 44, 48},
		},
		{
			[][]int{{1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1, 2}, {1}},
			[]int{1},
		},
		{
			[][]int{
				{97, 95, 26, 60, 35, 75, 83},
				{38, 68, 87, 75, 24, 51, 26, 27, 97, 48, 80, 37, 8, 98},
				{13, 18, 54, 40, 26, 100, 75, 5, 11, 97},
				{75, 33, 95, 7, 47, 23, 49, 26, 70},
				{4, 52, 26, 58, 93, 51, 75, 21, 14, 18},
				{89, 7, 18, 75, 56, 70, 31, 35, 37, 23, 26, 46, 65},
				{75, 98, 36, 26, 99, 51, 69, 94, 28, 43, 18},
				{43, 64, 76, 26, 39, 75, 30, 66, 45, 19},
				{35, 26, 75, 42, 36, 58},
			},
			[]int{26, 75},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := fucntion(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([][]int) []int
