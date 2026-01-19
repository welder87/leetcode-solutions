package problem3083

import (
	"testing"
)

func TestIsSubstringPresent(t *testing.T) {
	testIsSubstringPresent(t, isSubstringPresent)
}

func TestIsSubstringPresentV1(t *testing.T) {
	testIsSubstringPresent(t, isSubstringPresentV1)
}

func testIsSubstringPresent(t *testing.T, fn func(string) bool) {
	testCases := []struct {
		s    string
		ans  bool
		name string
	}{
		{"leetcode", true, "Preset case 1"},
		{"abcba", true, "Preset case 2"},
		{"abcd", false, "Preset case 3"},
		{"ausoee", true, "Preset case 4"},
		{"a", false, "Single character"},
		{"abcde", false, "No matching pairs"},
		{"aa", true, "Two same characters"},
		{"aaa", true, "Appears multiple times"},
		{"aba", true, "Palindrome odd length"},
		{"abba", true, "Palindrome even length"},
		{"ababa", true, "Palindrome"},
		{"aab", true, "Single pair in start"},
		{"baa", true, "Single pair in end"},
		{"abcdefghijklmnopqrstuvwxyz", false, "Long string no repeats"},
		{"abcdefedcba", true, "Long Palindrome"},
		{"abcabc", false, "First or second"},
		{"xyyx", true, "Pair in middle"},
		{"abccba", true, "Multiple valid pair"},
		{"abca", false, " No length-2 substring in reverse"},
		{repeatChar('a', 100), true, "All same character"},
		{repeatPattern("ab", 50), true, "Pattern ab"},
		{repeatPattern("abc", 33) + "ab", false, "Pattern with valid substring"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.s); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

func repeatChar(ch byte, n int) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = ch
	}
	return string(result)
}

func repeatPattern(pattern string, n int) string {
	result := make([]byte, 0, len(pattern)*n)
	for range n {
		result = append(result, pattern...)
	}
	return string(result)
}
