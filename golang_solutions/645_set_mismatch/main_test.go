package problem645

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	testFindMaxAverage(t, findErrorNums)
}

func TestFindMaxAverageV1(t *testing.T) {
	testFindMaxAverage(t, findErrorNumsV1)
}

func TestFindMaxAverageV2(t *testing.T) {
	testFindMaxAverage(t, findErrorNumsV2)
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
		{
			nums: []int{1, 2, 3, 4, 5, 6, 8, 8},
			ans:  []int{8, 7},
		},
		{
			nums: []int{8, 2, 4, 3, 7, 6, 5, 8},
			ans:  []int{8, 1},
		},
		{
			nums: []int{3, 2, 2},
			ans:  []int{2, 1},
		},
		{
			nums: []int{1, 7, 8, 4, 5, 6, 2, 8},
			ans:  []int{8, 3},
		},
		{
			nums: []int{3, 3, 2},
			ans:  []int{3, 1},
		},
		{
			nums: []int{3, 2, 1, 2},
			ans:  []int{2, 4},
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
