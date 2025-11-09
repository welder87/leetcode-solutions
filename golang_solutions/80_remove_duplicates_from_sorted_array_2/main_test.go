package problem80

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		expectedNums []int
		expectedAns  int
	}{
		// preset cases
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}, 7},
		// common cases
		{
			[]int{-1, -1, -1, -1, 0, 0, 0, 3, 3, 4},
			[]int{-1, -1, 0, 0, 3, 3, 4},
			7,
		},
		{
			[]int{-1, -1, -1, -1, 0, 0, 0, 3, 3, 4, 4, 4, 4},
			[]int{-1, -1, 0, 0, 3, 3, 4, 4},
			8,
		},
		{[]int{-1, 2, 2, 2, 2, 5}, []int{-1, 2, 2, 5}, 4},
		{[]int{-1, -1, -1, -1, -1}, []int{-1, -1}, 2},
		{[]int{-1, -1, 0, 0, 3, 3}, []int{-1, -1, 0, 0, 3, 3}, 6},
		// corner cases
		{[]int{-15}, []int{-15}, 1},
		{[]int{-15, -15}, []int{-15, -15}, 2},
		{[]int{0, 0, 0}, []int{0, 0}, 2},
		{[]int{-15, 0, 1, 4}, []int{-15, 0, 1, 4}, 4},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := removeDuplicates(testCase.nums)
			for i := 0; i < ans; i++ {
				if testCase.nums[i] != testCase.expectedNums[i] {
					t.Errorf(
						"got %v, want %v",
						testCase.nums,
						testCase.expectedNums,
					)
				}
			}
			if ans != testCase.expectedAns {
				t.Errorf("got %v, want %v", ans, testCase.expectedAns)
			}
		})
	}
}
