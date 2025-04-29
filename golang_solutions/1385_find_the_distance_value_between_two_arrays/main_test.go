package problem1385

import (
	"fmt"
	"reflect"
	"testing"
)

type fn func([]int, []int, int) int

func TestFindTheDistanceValue(t *testing.T) {
	testFindTheDistanceValue(t, findTheDistanceValue)
}

func testFindTheDistanceValue(t *testing.T, function fn) {
	testCases := []struct {
		arr1 []int
		arr2 []int
		d    int
		ans  int
	}{
		// preset cases
		{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
		// common cases

		// corner cases
		{[]int{1}, []int{-1}, 2, 0},
		{[]int{1}, []int{-1}, 1, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf(
			"%v %v %v", testCase.arr1, testCase.arr2, testCase.d,
		)
		t.Run(testName, func(t *testing.T) {
			result := function(testCase.arr1, testCase.arr2, testCase.d)
			if !reflect.DeepEqual(result, testCase.ans) {
				t.Errorf("got %v, want %v", result, testCase.ans)
			}
		})
	}
}
