package problem1508

import (
	"testing"
)

func TestRangeSum(t *testing.T) {
	testRangeSum(t, rangeSum)
}

func testRangeSum(t *testing.T, fn func([]int, int, int, int) int) {
	testCases := []struct {
		nums  []int
		n     int
		left  int
		right int
		ans   int
		name  string
	}{
		{[]int{1, 2, 3, 4}, 4, 1, 5, 13, "Preset Case 1"},
		{[]int{1, 2, 3, 4}, 4, 3, 4, 6, "Preset Case 2"},
		{[]int{1, 2, 3, 4}, 4, 1, 10, 50, "Preset Case 3"},
		{createRepeatedSlice(1000, 100), 1000, 1, 500500, 716699888, "Preset Case 4"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.nums, testCase.n, testCase.left, testCase.right)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

func createRepeatedSlice(n int, value int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = value
	}
	return s
}
