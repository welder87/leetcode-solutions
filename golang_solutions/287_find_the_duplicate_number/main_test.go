package problem287

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindDuplicate(t *testing.T) {
	testFindDuplicate(t, findDuplicate)
}

func TestFindDuplicateV1(t *testing.T) {
	testFindDuplicate(t, findDuplicateV1)
}

func TestFindDuplicateV2(t *testing.T) {
	testFindDuplicate(t, findDuplicateV2)
}

func testFindDuplicate(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
		// common cases
		{[]int{3, 3, 1, 2, 4}, 3},
		{[]int{3, 3, 3, 3, 1}, 3},
		{[]int{3, 2, 3, 3, 1}, 3},
		{[]int{3, 2, 3, 3, 1}, 3},
		// corner cases
		{[]int{1, 1}, 1},
		{[]int{1}, 0},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
