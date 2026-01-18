package problem383

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	testCanConstruct(t, canConstruct)
}

func testCanConstruct(t *testing.T, fn func(string, string) bool) {
	testCases := []struct {
		name       string
		ransomNote string
		magazine   string
		ans        bool
	}{
		{"Preset case 1", "a", "b", false},
		{"Preset case 2", "aa", "ab", false},
		{"Preset case 3", "ab", "aab", true},
		{"Valid construction - simple", "a", "ab", true},
		{"Valid construction - exact match", "aa", "aab", true},
		{"Invalid - insufficient letters", "aa", "ab", false},
		{"Invalid - missing letter", "abc", "ab", false},
		{"Valid - magazine has extra letters", "hello", "helloworld", true},
		{"Valid - same letters different order", "abc", "cba", true},
		{"Invalid - missing one instance", "aaa", "aa", false},
		{"Large ransom note valid", "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa", true},
		{"Large ransom note invalid", "aaaaaaaaaa", "aaaaaaaaa", false},
		{"Multiple same character needed", "aaaaab", "aaaaac", false},
		{"Multiple characters with correct frequency", "aabbcc", "aaabbbccc", true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.ransomNote, testCase.magazine)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
