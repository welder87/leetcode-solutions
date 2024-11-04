package problem2000

import (
	"reflect"
	"testing"
)

func TestReversePrefix(t *testing.T) {
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
		result := ReversePrefix(testCase.word, testCase.ch)
		t.Logf("Calling ReversePrefix(%v, %v)", testCase.word, string(testCase.ch))
		if !reflect.DeepEqual(testCase.expectedWord, result) {
			t.Errorf("Expected element type: %v, got: %v", testCase.expectedWord, result)
		}
	}
}
