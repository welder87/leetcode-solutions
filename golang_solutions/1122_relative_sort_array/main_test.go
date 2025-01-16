package problem1122

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	testRelativeSortArray(t, relativeSortArray)
}

func TestRelativeSortArrayV1(t *testing.T) {
	testRelativeSortArray(t, relativeSortArrayV1)
}

func testRelativeSortArray(t *testing.T, function fn) {
	testCases := []struct {
		arr1 []int
		arr2 []int
		ans  []int
	}{
		{
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			ans:  []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			arr1: []int{28, 6, 22, 8, 44, 17},
			arr2: []int{22, 28, 8, 6},
			ans:  []int{22, 28, 8, 6, 17, 44},
		},
		{
			arr1: []int{22},
			arr2: []int{22},
			ans:  []int{22},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.arr1, testCase.arr2)
		t.Run(testName, func(t *testing.T) {
			result := function(testCase.arr1, testCase.arr2)
			if !reflect.DeepEqual(result, testCase.ans) {
				t.Errorf("got %v, want %v", result, testCase.ans)
			}
		})
	}
}

type fn func([]int, []int) []int
