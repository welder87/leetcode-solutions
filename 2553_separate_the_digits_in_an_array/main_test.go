package problem2553

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSeparateDigits(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{7, 1, 3, 9},
			ans:  []int{7, 1, 3, 9},
		},
		{
			nums: []int{13, 25, 83, 77},
			ans:  []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		{
			nums: []int{100, 25},
			ans:  []int{1, 0, 0, 2, 5},
		},
		{
			nums: []int{},
			ans:  []int{},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := separateDigits(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
