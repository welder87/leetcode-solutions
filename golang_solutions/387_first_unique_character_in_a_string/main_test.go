package problem389

import (
	"strings"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	testFirstUniqChar(t, firstUniqChar)
}

func TestFirstUniqCharV1(t *testing.T) {
	testFirstUniqChar(t, firstUniqCharV1)
}

func testFirstUniqChar(t *testing.T, fn func(string) int) {
	testCases := []struct {
		strs string
		ans  int
		name string
	}{
		{"leetcode", 0, "Preset Case 1"},
		{"loveleetcode", 2, "Preset Case 2"},
		{"aabb", -1, "Preset Case 3"},
		{"a", 0, "Single character"},
		{"aa", -1, "Two same characters"},
		{"ab", 0, "Two different characters"},
		{"aaaaa", -1, "All same character"},
		{"abcde", 0, "First character is unique"},
		{"aabbc", 4, "Last character is unique"},
		{"aabcc", 2, "Middle character is unique"},
		{"abc", 0, "Multiple uniques, first one should be returned"},
		{"abcdefghijklmnopqrstuvwxyz", 0, "Long string with early unique"},
		{"aaaaaaaaaaaaaaaaaaaaab", 21, "Long string with late unique"},
		{"racecar", 3, "Palindrome with unique middle"},
		{"abccba", -1, "Palindrome no unique"},
		{"abcabcabcx", 9, "Repeating pattern with one unique"},
		{"aaabbbcccddde", 12, "Many duplicates then unique"},
		{"aabbcbbaa", 4, "Unique between duplicates"},
		{strings.Repeat("ababab", 100), -1, "Performance Case"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.strs); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
