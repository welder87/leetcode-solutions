package problem912

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	testSortArray(t, sortArray)
}

func TestSortArrayV1(t *testing.T) {
	testSortArray(t, sortArrayV1)
}

func testSortArray(t *testing.T, function fn) {
	testCases := []struct {
		name []int
		ans  []int
	}{
		{
			name: []int{5, 2, 3, 1},
			ans:  []int{1, 2, 3, 5},
		},
		{
			name: []int{5, 1, 1, 2, 0, 0},
			ans:  []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: []int{},
			ans:  []int{},
		},
		{
			name: []int{5, 4, 4, -25, 3, 2, -1, 1, -10},
			ans:  []int{-25, -10, -1, 1, 2, 3, 4, 4, 5},
		},
		{
			name: []int{1, -1, 0},
			ans:  []int{-1, 0, 1},
		},
		{
			name: []int{3, 5, 4},
			ans:  []int{3, 4, 5},
		},
		{
			name: []int{-3, -5, -4},
			ans:  []int{-5, -4, -3},
		},
		{
			name: []int{1, 1, 1},
			ans:  []int{1, 1, 1},
		},
		{
			name: []int{15000, 0, -30000},
			ans:  []int{-30000, 0, 15000},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.name)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.name)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) []int
