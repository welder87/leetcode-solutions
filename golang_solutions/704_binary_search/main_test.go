package problem704

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		ans    int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			ans:    4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			ans:    -1,
		},
		{
			nums:   []int{-3, -1, 3, 5, 9, 12},
			target: 12,
			ans:    5,
		},
		{
			nums:   []int{-3, -1, 3, 5, 9, 12},
			target: -1,
			ans:    1,
		},
		{
			nums:   []int{-3},
			target: -3,
			ans:    0,
		},
		{
			nums:   []int{-3},
			target: 1,
			ans:    -1,
		},
		{
			nums:   []int{},
			target: 25,
			ans:    -1,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := search(testCase.nums, testCase.target)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
