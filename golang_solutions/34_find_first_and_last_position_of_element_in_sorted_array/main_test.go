package problem34

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			ans:    []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{7, 7, 7, 7},
			target: 7,
			ans:    []int{0, 3},
		},
		{
			nums:   []int{7, 7, 7, 7},
			target: 6,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 8, 8, 10},
			target: 8,
			ans:    []int{3, 6},
		},
		{
			nums:   []int{5},
			target: 5,
			ans:    []int{0, 0},
		},
		{
			nums:   []int{1, 22, 22, 33, 50, 100, 20000},
			target: 33,
			ans:    []int{3, 3},
		},
		{
			nums:   []int{4, 6, 7, 7, 7, 20},
			target: 8,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{6, 7, 9, 10, 10, 10, 90},
			target: 10,
			ans:    []int{3, 5},
		},
		{
			nums:   []int{1, 3, 5, 8, 13, 21},
			target: 40,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{-25, -17, -15, 0, 3, 5, 5, 8, 13, 21},
			target: -45,
			ans:    []int{-1, -1},
		},
		{
			nums:   []int{-25, -17, -15, 0, 3, 5, 5, 8, 13, 21},
			target: 5,
			ans:    []int{5, 6},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := searchRange(testCase.nums, testCase.target)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
