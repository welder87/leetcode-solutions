package problem1394

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindLucky(t *testing.T) {
	testFindLucky(t, findLucky)
}

func TestFindLuckyV1(t *testing.T) {
	testFindLucky(t, findLuckyV1)
}

func testFindLucky(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{2, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{2, 2, 2, 3, 3}, -1},
		// common cases
		{[]int{1, 2, 2, 2, 3, 3, 3}, 3},
		{[]int{1, 4, 4, 4, 4, 3, 3, 3}, 4},
		// corner cases
		{[]int{1}, 1},
		{[]int{2}, -1},
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
