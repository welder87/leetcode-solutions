package problem680

import (
	"fmt"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		s   string
		ans bool
	}{
		{
			s:   "aba",
			ans: true,
		},
		{
			s:   "abca",
			ans: true,
		},
		{
			s:   "abc",
			ans: false,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.s)
		t.Run(testName, func(t *testing.T) {
			ans := validPalindrome(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
