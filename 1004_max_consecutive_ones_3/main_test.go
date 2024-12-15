package problem1004

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		ans  int
	}{
		{
			nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
			ans:  6,
		},
		{
			nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
			ans:  10,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := longestOnes(testCase.nums, testCase.k)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
