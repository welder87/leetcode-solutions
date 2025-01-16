package problem1876

import (
	"fmt"
	"testing"
)

func TestCountGoodSubstrings(t *testing.T) {
	testCases := []struct {
		s   string
		ans int
	}{
		{
			s:   "xyzzaz",
			ans: 1,
		},
		{
			s:   "aababcabc",
			ans: 4,
		},
		{
			s:   "xy",
			ans: 0,
		},
		{
			s:   "xyz",
			ans: 1,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.s)
		t.Run(testName, func(t *testing.T) {
			ans := countGoodSubstrings(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
