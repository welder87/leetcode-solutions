package problem645

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	testFindMaxAverage(t, findErrorNums)
}

func testFindMaxAverage(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{1, 2, 2, 4},
			ans:  []int{2, 3},
		},
		{
			nums: []int{1, 1},
			ans:  []int{1, 2},
		},
		{
			nums: []int{1, 2, 3, 3},
			ans:  []int{3, 4},
		},
		{
			nums: []int{2, 2},
			ans:  []int{2, 1},
		},
		{
			nums: []int{1, 2, 8, 4, 5, 6, 7, 8},
			ans:  []int{8, 3},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) []int
