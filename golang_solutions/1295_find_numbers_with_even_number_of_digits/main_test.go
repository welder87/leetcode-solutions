package problem1295

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindNumbers(t *testing.T) {
	testFindNumbers(t, findNumbers)
}

func TestFindNumbersV1(t *testing.T) {
	testFindNumbers(t, findNumbersV1)
}

func testFindNumbers(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
		// common cases
		// corner cases
		{[]int{0}, 0},
		{[]int{99}, 1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
