package problem2000

import (
	"reflect"
	"testing"
)

func TestReversePrefix(t *testing.T) {
	testReversePrefix(t, ReversePrefix)
}

func TestReversePrefixV1(t *testing.T) {
	testReversePrefix(t, ReversePrefixV1)
}

type fn func(string, byte) string

func testReversePrefix(t *testing.T, function fn) {
	for _, testCase := range getTestCases() {
		result := function(testCase.word, testCase.ch)
		t.Logf("Calling ReversePrefix(%v, %v)", testCase.word, string(testCase.ch))
		if !reflect.DeepEqual(testCase.expectedWord, result) {
			t.Errorf("Expected element type: %v, got: %v", testCase.expectedWord, result)
		}
	}
}

type testCase struct {
	word         string
	ch           byte
	expectedWord string
}

func getTestCases() []testCase {
	return []testCase{
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
}
