package problem1805

import (
	"fmt"
	"testing"
)

func TestNumDifferentIntegers(t *testing.T) {
	testCases := []struct {
		word string
		ans  int
	}{
		{
			word: "a123bc34d8ef34",
			ans:  3,
		},
		{
			word: "leet1234code234",
			ans:  2,
		},
		{
			word: "a1b01c001",
			ans:  1,
		},
		{
			word: "0leet1234code234do1",
			ans:  4,
		},
		{
			word: "au35cf7c305",
			ans:  3,
		},
		{
			word: "00123456",
			ans:  1,
		},
		{
			word: "",
			ans:  0,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.word)
		t.Run(testName, func(t *testing.T) {
			ans := numDifferentIntegers(testCase.word)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
