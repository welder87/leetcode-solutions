package problem20

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s   string
		ans bool
	}{
		{
			s:   "()",
			ans: true,
		},
		{
			s:   "()[]{}",
			ans: true,
		},
		{
			s:   "(]",
			ans: false,
		},
		{
			s:   "([])",
			ans: true,
		},
		{
			s:   "[({)]",
			ans: false,
		},
		{
			s:   "[({}[()])]",
			ans: true,
		},
		{
			s:   "",
			ans: true,
		},
		{
			s:   "[({}[(})])]",
			ans: false,
		},
		{
			s:   "])]",
			ans: false,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.s)
		t.Run(testName, func(t *testing.T) {
			ans := isValid(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
