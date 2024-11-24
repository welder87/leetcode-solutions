package problem1446

import (
	"fmt"
	"testing"
)

func TestMaxPower(t *testing.T) {
	testCases := []struct {
		s   string
		ans int
	}{
		{
			s:   "leetcode",
			ans: 2,
		},
		{
			s:   "abbcccddddeeeeedcba",
			ans: 5,
		},
		{
			s:   "abbcccddddeeeeedceeeeeba",
			ans: 5,
		},
		{
			s:   "abbcccddddeeeee",
			ans: 5,
		},
		{
			s:   "eeeeeabbcccdddd",
			ans: 5,
		},
		{
			s:   "eeeee",
			ans: 5,
		},
		{
			s:   "e",
			ans: 1,
		},
		{
			s:   "abbcccddddeeeeedckkkkkba",
			ans: 5,
		},
		{
			s:   "abcde",
			ans: 1,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.s)
		t.Run(testName, func(t *testing.T) {
			ans := maxPower(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
