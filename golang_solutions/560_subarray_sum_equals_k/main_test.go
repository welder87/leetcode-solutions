package problem560

import (
	"fmt"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	testSubarraySum(t, subarraySum)
}

func testSubarraySum(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		k    int
		ans  int
	}{
		{nums: []int{1, 1, 1}, k: 2, ans: 2},
		{nums: []int{1, 2, 3}, k: 3, ans: 2},
		{nums: []int{1}, k: 1, ans: 1},
		{nums: []int{1}, k: -1, ans: 0},
		{nums: []int{-5, 7, 3, -2, 4, -2, 12}, k: 12, ans: 3},
		{nums: []int{1, -1, 0}, k: 0, ans: 3},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums, testCase.k)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int, int) int
