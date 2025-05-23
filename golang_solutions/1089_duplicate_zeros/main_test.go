package problem1086

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	testCases := []struct {
		nums         []int
		expectedNums []int
	}{
		{
			nums:         []int{1, 0, 2, 3, 0, 4, 5, 0},
			expectedNums: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			nums:         []int{1, 2, 3},
			expectedNums: []int{1, 2, 3},
		},
		{
			nums:         []int{0, 2, 1, 3, 0, 4, 5, 0},
			expectedNums: []int{0, 0, 2, 1, 3, 0, 0, 4},
		},
		{
			nums:         []int{2, 1, 3, 0, 0, 4, 5, 3},
			expectedNums: []int{2, 1, 3, 0, 0, 0, 0, 4},
		},
		{
			nums:         []int{2, 0, 3, 0, 1, 0, 5, 3},
			expectedNums: []int{2, 0, 0, 3, 0, 0, 1, 0},
		},
		{
			nums:         []int{1, 5, 2, 0, 6, 8, 0, 6, 0},
			expectedNums: []int{1, 5, 2, 0, 0, 6, 8, 0, 0},
		},
		{
			nums:         []int{9, 9, 9, 4, 8, 0, 0, 3, 7, 2, 0, 0, 0, 0, 9, 1, 0, 0, 1, 1, 0, 5, 6, 3, 1, 6, 0, 0, 2, 3, 4, 7, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 7, 3, 3, 0, 5, 7, 0, 8, 1, 9, 6, 3, 0, 8, 8, 8, 8, 0, 0, 5, 0, 0, 0, 3, 7, 7, 7, 7, 5, 1, 0, 0, 8, 0, 0},
			expectedNums: []int{9, 9, 9, 4, 8, 0, 0, 0, 0, 3, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 1, 1, 0, 0, 5, 6, 3, 1, 6, 0, 0, 0, 0, 2, 3, 4, 7, 0, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 0, 0, 7, 3, 3, 0, 0, 5, 7, 0, 0, 8, 1, 9, 6, 3, 0, 0, 8, 8, 8, 8, 0},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			DuplicateZeros(testCase.nums)
			if !reflect.DeepEqual(testCase.expectedNums, testCase.nums) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.expectedNums)
			}
		})
	}
}
