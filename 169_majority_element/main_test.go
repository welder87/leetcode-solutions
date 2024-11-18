package problem169

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testMajorityElement(t, majorityElement)
}

func testMajorityElement(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 2, 3},
			ans:  3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			ans:  2,
		},
		{
			nums: []int{2},
			ans:  2,
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
