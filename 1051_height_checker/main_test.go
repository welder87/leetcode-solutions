package problem1051

import (
	"fmt"
	"testing"
)

func TestHeightChecker(t *testing.T) {
	testCases := []struct {
		heights []int
		ans     int
	}{
		{
			heights: []int{1, 1, 4, 2, 1, 3},
			ans:     3,
		},
		{
			heights: []int{5, 1, 2, 3, 4},
			ans:     5,
		},
		{
			heights: []int{1, 2, 3, 4, 5},
			ans:     0,
		},
		{
			heights: []int{1},
			ans:     0,
		},
		{
			heights: []int{5, 4, 3, 2, 1},
			ans:     4,
		},
		{
			heights: []int{5, 4, 1, 2, 3},
			ans:     5,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.heights)
		t.Run(testName, func(t *testing.T) {
			ans := heightChecker(testCase.heights)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
