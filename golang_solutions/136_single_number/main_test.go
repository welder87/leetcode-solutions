package problem136

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testSingleNumber(t, singleNumber)
}

func TestSingleNumberV1(t *testing.T) {
	testSingleNumber(t, singleNumberV1)
}

func testSingleNumber(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{2, 2, 1},
			ans:  1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			ans:  4,
		},
		{
			nums: []int{1},
			ans:  1,
		},
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

type fn func([]int) int
