package problem35

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		ans    int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			ans:    2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			ans:    1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			ans:    4,
		},
		{
			nums:   []int{1},
			target: 2,
			ans:    1,
		},
		{
			nums:   []int{1},
			target: -1,
			ans:    0,
		},
		{
			nums:   []int{-25, -15, -6, 0, 1, 7, 10, 16},
			target: 3,
			ans:    5,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := searchInsert(testCase.nums, testCase.target)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
