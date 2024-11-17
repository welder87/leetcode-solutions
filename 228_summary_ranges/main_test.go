package problem228

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			ans:  []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			ans:  []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums: []int{-10, -5, -4, -3, 0, 1, 2, 3, 4, 5},
			ans:  []string{"-10", "-5->-3", "0->5"},
		},
		{
			nums: []int{-10, -5, -1, 1, 3, 5, 8, 10},
			ans:  []string{"-10", "-5", "-1", "1", "3", "5", "8", "10"},
		},
		{
			nums: []int{-10, -9, -8, -7, -6, -5},
			ans:  []string{"-10->-5"},
		},
		{
			nums: []int{-10, -9, -8, -7, -6, -5, 0},
			ans:  []string{"-10->-5", "0"},
		},
		{
			nums: []int{},
			ans:  []string{},
		},
		{
			nums: []int{1},
			ans:  []string{"1"},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := summaryRanges(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
