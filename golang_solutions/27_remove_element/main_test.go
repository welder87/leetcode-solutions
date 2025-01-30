package problem27

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type fn func([]int, int) int

func TestRemoveElement(t *testing.T) {
	tesRemoveElement(t, removeElement)
}

func TestRemoveElementV1(t *testing.T) {
	tesRemoveElement(t, removeElementV1)
}

func tesRemoveElement(t *testing.T, function fn) {
	testCases := []struct {
		nums         []int
		val          int
		ans          int
		expectedNums []int
	}{
		{
			nums:         []int{3, 2, 2, 3},
			val:          3,
			ans:          2,
			expectedNums: []int{2, 2, -1, -1},
		},
		{
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			ans:          5,
			expectedNums: []int{0, 1, 4, 0, 3, -1, -1, -1},
		},
		{
			nums:         []int{2, 1, 2, 2, 3, 0, 4, 2},
			val:          2,
			ans:          4,
			expectedNums: []int{4, 1, 0, 3, -1, -1, -1, -1},
		},
		{
			nums:         []int{},
			val:          2,
			ans:          0,
			expectedNums: []int{},
		},
		{
			nums:         []int{1},
			val:          2,
			ans:          1,
			expectedNums: []int{1},
		},
		{
			nums:         []int{2},
			val:          2,
			ans:          0,
			expectedNums: []int{-1},
		},
		{
			nums:         []int{2},
			val:          2,
			ans:          0,
			expectedNums: []int{-1},
		},
		{
			nums:         []int{2, 1, 0, 2, 3, 0, 4, 2},
			val:          0,
			ans:          6,
			expectedNums: []int{2, 1, 2, 2, 3, 4, -1, -1},
		},
		{
			nums:         []int{2, 2, 2},
			val:          2,
			ans:          0,
			expectedNums: []int{-1, -1, -1},
		},
		{
			nums:         []int{0, 1, 0, 2, 3, 0, 4, 1},
			val:          2,
			ans:          7,
			expectedNums: []int{0, 1, 0, 1, 3, 0, 4, -1},
		},
		{
			nums:         []int{0, 1, 0, 2, 3, 0, 4, 1},
			val:          5,
			ans:          8,
			expectedNums: []int{0, 1, 0, 2, 3, 0, 4, 1},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := removeElement(testCase.nums, testCase.val)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
			sort.Ints(testCase.nums)
			sort.Ints(testCase.expectedNums)
			if !reflect.DeepEqual(testCase.nums, testCase.expectedNums) {
				t.Errorf("got %v, want %v", testCase.nums, testCase.expectedNums)
			}
		})
	}
}
