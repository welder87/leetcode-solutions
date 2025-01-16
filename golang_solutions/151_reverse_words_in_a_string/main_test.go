package problem151

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	testReverseWords(t, reverseWords)
}

func testReverseWords(t *testing.T, function fn) {
	testCases := []struct {
		s   string
		ans string
	}{
		{
			s:   "the sky is blue",
			ans: "blue is sky the",
		},
		{
			s:   "  hello world  ",
			ans: "world hello",
		},
		{
			s:   "a good   example",
			ans: "example good a",
		},
		{
			s:   "a    good   example  a",
			ans: "a example good a",
		},
		{
			s:   "  a    good   example  a  ",
			ans: "a example good a",
		},
		{
			s:   " a ",
			ans: "a",
		},
		{
			s:   "a",
			ans: "a",
		},
		{
			s:   "a b",
			ans: "b a",
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

type fn func(string) string
