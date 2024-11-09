package problem219

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	testCases := []struct {
		nums        []int
		k           int
		expectedAns bool
	}{
		{
			nums:        []int{1, 2, 3, 1},
			k:           3,
			expectedAns: true,
		},
		{
			nums:        []int{1, 0, 1, 1},
			k:           1,
			expectedAns: true,
		},
		{
			nums:        []int{1, 2, 3, 1, 2, 3},
			k:           2,
			expectedAns: false,
		},
		{
			nums:        []int{-1, -1, -1, -1, -1},
			k:           0,
			expectedAns: false,
		},
		{
			nums:        []int{},
			k:           10000,
			expectedAns: false,
		},
		{
			nums:        []int{1},
			k:           10000,
			expectedAns: false,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := containsNearbyDuplicate(testCase.nums, testCase.k)
			if !reflect.DeepEqual(testCase.expectedAns, ans) {
				t.Errorf("got %v, want %v", ans, testCase.expectedAns)
			}
		})
	}
}
