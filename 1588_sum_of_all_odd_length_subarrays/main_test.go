package problem1588

import (
	"fmt"
	"testing"
)

func TestSumOddLengthSubarrays(t *testing.T) {
	testSumOddLengthSubarrays(t, sumOddLengthSubarrays)
}

func testSumOddLengthSubarrays(t *testing.T, function fn) {
	testCases := []struct {
		s   []int
		ans int
	}{
		{
			s:   []int{1, 4, 2, 5, 3},
			ans: 58,
		},
		{
			s:   []int{1, 2},
			ans: 3,
		},
		{
			s:   []int{10, 11, 12},
			ans: 66,
		},
		{
			s:   []int{10},
			ans: 10,
		},
		{
			s:   []int{1, 4, 2, 5, 3, 6},
			ans: 98,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.s)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) int
