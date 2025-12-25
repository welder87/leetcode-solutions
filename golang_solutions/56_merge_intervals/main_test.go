package problem56

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testMerge(t, merge)
}

func TestMergeV1(t *testing.T) {
	testMerge(t, mergeV1)
}

func TestMergeV2(t *testing.T) {
	testMerge(t, mergeV2)
}

func testMerge(t *testing.T, fn func([][]int) [][]int) {
	testCases := []struct {
		name      string
		intervals [][]int
		ans       [][]int
	}{
		{
			"Preset case 1 (overlapping intervals)",
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"Preset case 2 (adjacent intervals)",
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			"Preset case 3 (unsorted overlapping)",
			[][]int{{4, 7}, {1, 4}},
			[][]int{{1, 7}},
		},
		{
			"Sorted overlapping",
			[][]int{{1, 1}, {2, 3}, {2, 6}, {2, 15}},
			[][]int{{1, 1}, {2, 15}},
		},
		{
			"Single interval",
			[][]int{{4, 7}},
			[][]int{{4, 7}},
		},
		{
			"No overlaps",
			[][]int{{1, 2}, {3, 6}, {7, 15}},
			[][]int{{1, 2}, {3, 6}, {7, 15}},
		},
		{
			"All overlapping (first)",
			[][]int{{1, 4}, {2, 3}, {0, 5}},
			[][]int{{0, 5}},
		},
		{
			"All overlapping (second)",
			[][]int{{1, 3}, {2, 6}, {5, 15}, {15, 18}},
			[][]int{{1, 18}},
		},
		{
			"Unsorted nested intervals (first)",
			[][]int{{1, 4}, {3, 5}, {2, 4}, {7, 10}},
			[][]int{{1, 5}, {7, 10}},
		},
		{
			"Unsorted nested intervals (second)",
			[][]int{{7, 15}, {1, 1}, {1, 6}, {15, 15}},
			[][]int{{1, 6}, {7, 15}},
		},
		{
			"Touching intervals multiple",
			[][]int{{1, 2}, {2, 3}, {3, 4}},
			[][]int{{1, 4}},
		},
		{
			"Unsorted with gaps",
			[][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}},
			[][]int{{1, 3}, {4, 7}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.intervals)

			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, ans %v", ans, testCase.ans)
			}
		})
	}
}
