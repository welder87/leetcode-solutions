package problem448

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	testFindDisappearedNumbers(t, findDisappearedNumbers)
}

func testFindDisappearedNumbers(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			ans:  []int{5, 6},
		},
		{
			nums: []int{1, 1},
			ans:  []int{2},
		},
		{
			nums: []int{1, 3, 3},
			ans:  []int{2},
		},
		{
			nums: []int{2, 2, 2, 2, 2, 2, 2, 2},
			ans:  []int{1, 3, 4, 5, 6, 7, 8},
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
