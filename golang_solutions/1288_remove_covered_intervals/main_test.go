package problem1288

import (
	"testing"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	testRemoveCoveredIntervals(t, removeCoveredIntervals)
}

func testRemoveCoveredIntervals(t *testing.T, fn func([][]int) int) {
	testCases := []struct {
		name      string
		intervals [][]int
		ans       int
	}{
		{"Preset case 1", [][]int{{1, 4}, {3, 6}, {2, 8}}, 2},
		{"Preset case 2", [][]int{{1, 4}, {2, 3}}, 1},
		{"Not covered", [][]int{{6, 10}, {1, 4}, {3, 6}}, 3},
		{"All covered", [][]int{{1, 4}, {1, 10}, {3, 6}, {7, 8}}, 1},
		{"Single interval", [][]int{{1, 5}}, 1},
		{"Two identical intervals", [][]int{{1, 3}, {1, 3}}, 1},
		{"Non-overlapping intervals", [][]int{{1, 2}, {3, 4}, {5, 6}}, 3},
		{"Non-overlapping but touching", [][]int{{1, 2}, {2, 3}, {3, 4}}, 3},
		{"One interval covers all others", [][]int{{2, 5}, {3, 6}, {1, 10}, {4, 7}}, 1},
		{"Nested intervals", [][]int{{2, 9}, {3, 8}, {1, 10}, {4, 7}}, 1},
		{"Partially overlapping intervals", [][]int{{1, 3}, {2, 5}, {4, 6}}, 3},
		{"Overlap at one point", [][]int{{3, 4}, {1, 2}, {2, 3}}, 3},
		{"Same start, different ends", [][]int{{1, 4}, {1, 3}, {1, 5}}, 1},
		{"Same start, one covers other", [][]int{{1, 5}, {1, 3}}, 1},
		{"Same end, different starts", [][]int{{1, 5}, {2, 5}, {3, 5}}, 1},
		{"Same end, one covers other", [][]int{{1, 5}, {3, 5}}, 1},
		{"Mixed coverage", [][]int{{1, 5}, {2, 4}, {3, 6}, {7, 9}}, 3},
		{"Complex case 1", [][]int{{1, 4}, {1, 5}, {2, 6}, {3, 7}, {8, 9}}, 4},
		{"Complex case 2", [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 6}}, 3},
		{"Large range intervals", [][]int{{0, 100}, {10, 20}, {30, 40}, {50, 60}}, 1},
		{"Very large numbers", [][]int{{1, 100_000}, {50_000, 60_000}, {2, 99_999}}, 1},
		{"Unsorted intervals", [][]int{{3, 6}, {1, 4}, {2, 8}}, 2},
		{"Reverse sorted intervals", [][]int{{10, 12}, {5, 8}, {1, 4}}, 3},
		{"Zero-length intervals", [][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{"Zero-length covered by regular", [][]int{{1, 1}, {1, 5}}, 1},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.intervals)
			if ans != testCase.ans {
				t.Errorf("got %v, ans %v", ans, testCase.ans)
			}
		})
	}
}
