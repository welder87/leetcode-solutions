package problem1047

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicates)
}

func TestRemoveDuplicatesV1(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicatesV1)
}

func TestRemoveDuplicatesV2(t *testing.T) {
	testRemoveDuplicates(t, removeDuplicatesV2)
}

func testRemoveDuplicates(t *testing.T, fn func(string) string) {
	testCases := []struct {
		name string
		s    string
		ans  string
	}{
		{"Preset Case 1", "abbaca", "ca"},
		{"Preset Case 2", "azxxzy", "ay"},
		{"No duplicates", "abcdef", "abcdef"},
		{"Even duplicate characters ", "aaaa", ""},
		{"Odd duplicate characters ", "aaaaa", "a"},
		{"Single character", "a", "a"},
		{"Nested duplicates", "abba", ""},
		{"Alternating characters 1", "aba", "aba"},
		{"Alternating characters 2", "abab", "abab"},
		{"First unique", "caa", "c"},
		{"Last unique", "aac", "c"},
		{"Complex case 1", "abbacaabbaca", "caca"},
		{"Complex case 1", "abbacaabbddddca", "a"},
		{"Complex case 2", "abccbad", "d"},
		{"Complex case 3", "aabbvvv", "v"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.s); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
