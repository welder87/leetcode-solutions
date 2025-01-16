package problem80

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		expectedNums []int
		expectedAns  int
	}{
		{
			nums:         []int{1, 1, 1, 2, 2, 3},
			expectedNums: []int{1, 1, 2, 2, 3},
			expectedAns:  5,
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
			expectedAns:  7,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := removeDuplicates(testCase.nums)
			if !reflect.DeepEqual(testCase.expectedNums, testCase.nums) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.expectedNums)
			}
			if ans != testCase.expectedAns {
				t.Errorf("got %v, want %v", ans, testCase.expectedAns)
			}
		})
	}
}
