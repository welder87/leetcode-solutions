package problem986

import (
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	testIntervalIntersection(t, intervalIntersection)
}

func testIntervalIntersection(t *testing.T, fn func([][]int, [][]int) [][]int) {
	testCases := []struct {
		name       string
		firstList  [][]int
		secondList [][]int
		ans        [][]int
	}{
		{
			"Preset case 1",
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			"Preset case 2",
			[][]int{{1, 3}, {5, 9}},
			[][]int{},
			[][]int{},
		},
		{
			"complex case",
			[][]int{{3, 5}, {9, 20}},
			[][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}},
			[][]int{{4, 5}, {9, 10}, {11, 12}, {14, 15}, {16, 20}},
		},
		{
			"nested within one (first)",
			[][]int{{0, 30}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		},
		{
			"nested within one (second)",
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{0, 30}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
		},
		{
			"one inside another (first)",
			[][]int{{1, 10}},
			[][]int{{3, 7}},
			[][]int{{3, 7}},
		},
		{
			"one inside another (second)",
			[][]int{{3, 7}},
			[][]int{{1, 10}},
			[][]int{{3, 7}},
		},
		{
			"ending at same point",
			[][]int{{1, 4}, {6, 8}},
			[][]int{{2, 4}, {5, 8}},
			[][]int{{2, 4}, {6, 8}},
		},
		{
			"starting at same point",
			[][]int{{1, 3}, {5, 7}},
			[][]int{{1, 2}, {5, 6}},
			[][]int{{1, 2}, {5, 6}},
		},
		{
			"touch at point (first-second)",
			[][]int{{0, 2}, {5, 10}},
			[][]int{{10, 24}, {15, 24}},
			[][]int{{10, 10}},
		},
		{
			"touch at point (second-first)",
			[][]int{{10, 24}, {15, 24}},
			[][]int{{0, 2}, {5, 10}},
			[][]int{{10, 10}},
		},
		{
			"touch at point (middle)",
			[][]int{{1, 3}, {4, 6}},
			[][]int{{3, 4}},
			[][]int{{3, 3}, {4, 4}},
		},
		{
			"Partial overlap of two (first)",
			[][]int{{24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{24, 24}, {25, 25}},
		},
		{
			"Partial overlap of two (second)",
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{24, 25}},
			[][]int{{24, 24}, {25, 25}},
		},
		{
			"partial overlap right",
			[][]int{{1, 4}},
			[][]int{{3, 6}},
			[][]int{{3, 4}},
		},
		{
			"partial overlap left",
			[][]int{{3, 6}},
			[][]int{{1, 4}},
			[][]int{{3, 4}},
		},
		{
			"non-overlapping intervals",
			[][]int{{1, 3}, {7, 9}},
			[][]int{{4, 6}, {10, 12}},
			[][]int{},
		},
		{
			"non-overlapping intervals (first)",
			[][]int{{0, 2}, {5, 10}},
			[][]int{{15, 24}, {15, 24}},
			[][]int{},
		},
		{
			"non-overlapping intervals (second)",
			[][]int{{15, 24}, {15, 24}},
			[][]int{{0, 2}, {5, 10}},
			[][]int{},
		},
		{
			"disjoint point",
			[][]int{{0, 0}},
			[][]int{{1, 1}},
			[][]int{},
		},
		{
			"identical intervals",
			[][]int{{1, 3}, {4, 6}},
			[][]int{{1, 3}, {4, 6}},
			[][]int{{1, 3}, {4, 6}},
		},
		{
			"empty first",
			[][]int{},
			[][]int{{1, 3}, {5, 9}},
			[][]int{},
		},
		{
			"empty second",
			[][]int{{1, 3}, {5, 9}},
			[][]int{},
			[][]int{},
		},
		{
			"both empty",
			[][]int{},
			[][]int{},
			[][]int{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.firstList, testCase.secondList)

			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, ans %v", testCase.firstList, testCase.secondList)
			}
		})
	}
}
