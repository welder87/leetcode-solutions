package problem1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testTwoSum(t, twoSum)
}

func TestTwoSumV1(t *testing.T) {
	testTwoSum(t, twoSumV1)
}

func testTwoSum(t *testing.T, function fn) {
	testCases := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			ans:    []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			ans:    []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			ans:    []int{0, 1},
		},
		// yandex algorithms interview
		{
			nums:   []int{1, -2, 4, 7, 3},
			target: 5,
			ans:    []int{0, 2},
		},
		{
			nums:   []int{0, 2, -4, 4, 6, 11, -7},
			target: 7,
			ans:    []int{2, 5},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums, testCase.target)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int, int) []int
