package problem57

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testInsert(t, insert)
}

func testInsert(t *testing.T, fn func([][]int, []int) [][]int) {
	testCases := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		ans         [][]int
	}{
		{
			"Preset case 1",
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
		},
		{
			"Preset case 2",
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			"no overlapping (middle)",
			[][]int{{1, 2}, {5, 6}, {10, 15}, {17, 18}},
			[]int{3, 4},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {10, 15}, {17, 18}},
		},
		{
			"no overlapping (end)",
			[][]int{{1, 2}, {5, 6}, {10, 15}, {17, 18}},
			[]int{19, 21},
			[][]int{{1, 2}, {5, 6}, {10, 15}, {17, 18}, {19, 21}},
		},
		{
			"no overlapping (start)",
			[][]int{{3, 4}, {5, 6}, {10, 15}},
			[]int{1, 2},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {10, 15}},
		},
		{
			"all overlapping",
			[][]int{{3, 4}, {5, 6}, {10, 15}},
			[]int{1, 17},
			[][]int{{1, 17}},
		},
		{
			"overlapping (start left)",
			[][]int{{3, 4}, {5, 6}, {10, 15}},
			[]int{1, 4},
			[][]int{{1, 4}, {5, 6}, {10, 15}},
		},
		{
			"overlapping (start right)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{2, 4},
			[][]int{{1, 4}, {5, 6}, {10, 15}},
		},
		{
			"fully covered (start)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{1, 3},
			[][]int{{1, 3}, {5, 6}, {10, 15}},
		},
		{
			"partially covered (start right)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{1, 2},
			[][]int{{1, 3}, {5, 6}, {10, 15}},
		},
		{
			"partially covered (start left)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{0, 3},
			[][]int{{0, 3}, {5, 6}, {10, 15}},
		},
		{
			"overlapping (end left)",
			[][]int{{3, 4}, {5, 6}, {10, 15}},
			[]int{9, 14},
			[][]int{{3, 4}, {5, 6}, {9, 15}},
		},
		{
			"overlapping (end right)",
			[][]int{{3, 4}, {5, 6}, {10, 15}},
			[]int{11, 16},
			[][]int{{3, 4}, {5, 6}, {10, 16}},
		},
		{
			"fully covered (end)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{10, 15},
			[][]int{{1, 3}, {5, 6}, {10, 15}},
		},
		{
			"partially covered (end right)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{10, 14},
			[][]int{{1, 3}, {5, 6}, {10, 15}},
		},
		{
			"partially covered (end left)",
			[][]int{{1, 3}, {5, 6}, {10, 15}},
			[]int{11, 15},
			[][]int{{1, 3}, {5, 6}, {10, 15}},
		},
		{
			"overlapping multiple",
			[][]int{{1, 2}, {4, 5}, {6, 7}, {8, 9}, {12, 16}},
			[]int{3, 10},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.intervals, testCase.newInterval)
			if !reflect.DeepEqual(ans, testCase.ans) {
				t.Errorf("got %v, ans %v", ans, testCase.ans)
			}
		})
	}
}
