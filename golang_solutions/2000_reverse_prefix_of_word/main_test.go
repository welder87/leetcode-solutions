package problem2000

import (
	"fmt"
	"testing"
)

func TestReversePrefix(t *testing.T) {
	testReversePrefix(t, ReversePrefix)
}

func TestReversePrefixV1(t *testing.T) {
	testReversePrefix(t, ReversePrefixV1)
}

func testReversePrefix(t *testing.T, function fn) {
	testCases := []struct {
		word         string
		ch           byte
		expectedWord string
	}{
		{
			word:         "abcdefd",
			ch:           byte('d'),
			expectedWord: "dcbaefd",
		},
		{
			word:         "xyxzxe",
			ch:           byte('z'),
			expectedWord: "zxyxxe",
		},
		{
			word:         "abcd",
			ch:           byte('z'),
			expectedWord: "abcd",
		},
		{
			word:         "abcvefd",
			ch:           byte('d'),
			expectedWord: "dfevcba",
		},
		{
			word:         "abcvefd",
			ch:           byte('f'),
			expectedWord: "fevcbad",
		},
		{
			word:         "akcvefd",
			ch:           byte('a'),
			expectedWord: "akcvefd",
		},
		{
			word:         "akcvefd",
			ch:           byte('k'),
			expectedWord: "kacvefd",
		},
		{
			word:         "",
			ch:           byte('k'),
			expectedWord: "",
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.word, testCase.ch)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.word, testCase.ch)
			if testCase.expectedWord != ans {
				t.Errorf("got %v, want %v", ans, testCase.expectedWord)
			}
		})
	}
}

type fn func(string, byte) string
