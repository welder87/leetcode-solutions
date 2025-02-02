package problem283

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{}, []int{}},
		{[]int{1, 3, 12, 0, 0}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 1, 15, 0, 12, 3, 15}, []int{1, 15, 12, 3, 15, 0, 0}},
		{
			[]int{0, 0, 0, -15, 10, -1, 0, 12, 0, 3, 15},
			[]int{-15, 10, -1, 12, 3, 15, 0, 0, 0, 0, 0},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			moveZeroes(testCase.nums)
			if !reflect.DeepEqual(testCase.nums, testCase.ans) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.ans)
			}
		})
	}
}
