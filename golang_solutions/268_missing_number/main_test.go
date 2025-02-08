package problem268

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestMissingNumber(t *testing.T) {
	testMissingNumber(t, missingNumber)
}

func TestMissingNumberV1(t *testing.T) {
	testMissingNumber(t, missingNumberV1)
}

func TestMissingNumberV2(t *testing.T) {
	testMissingNumber(t, missingNumberV2)
}

func TestMissingNumberV3(t *testing.T) {
	testMissingNumber(t, missingNumberV3)
}

func TestMissingNumberV4(t *testing.T) {
	testMissingNumber(t, missingNumberV4)
}

func testMissingNumber(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		// preset cases
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{1, 8, 4, 6, 2, 0, 9, 7, 5}, 3},
		// common cases
		{[]int{9, 6, 4, 2, 3, 5, 7, 8, 1}, 0},
		{[]int{9, 6, 4, 2, 3, 5, 7, 8, 0, 1}, 10},
		{[]int{2, 0, 1}, 3},
		{[]int{3, 0, 1, 5, 4}, 2},
		{[]int{3, 0, 4, 5, 2}, 1},
		{[]int{2, 3, 4, 5, 1}, 0},
		{[]int{2, 3, 0, 5, 1}, 4},
		// corner cases
		{[]int{1}, 0},
		{[]int{0}, 1},
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
