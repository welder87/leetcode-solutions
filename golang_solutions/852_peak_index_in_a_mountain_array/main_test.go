package problem852

import (
	"fmt"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		// common cases
		{[]int{5, 6, 4, 3, 2, 1, 0}, 1},
		{[]int{5, 6, 15, 25, 31, 7}, 4},
		{[]int{5, 6, 15, 14, 11, 0}, 2},
		// corner cases
		{[]int{7, 8, 4}, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := peakIndexInMountainArray(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
