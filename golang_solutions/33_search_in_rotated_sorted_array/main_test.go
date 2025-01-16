package problem33

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
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			ans:    4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			ans:    -1,
		},
		{
			nums:   []int{1},
			target: 1,
			ans:    0,
		},
		{
			nums:   []int{1},
			target: 25,
			ans:    -1,
		},
		{
			nums:   []int{2, 4, 5, 6, 7, 0, 1},
			target: 0,
			ans:    5,
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
