package problem2716

import (
	"testing"
)

func TestMinimizedStringLength(t *testing.T) {
	testMinimizedStringLength(t, minimizedStringLength)
}
func TestMinimizedStringLengthV1(t *testing.T) {
	testMinimizedStringLength(t, minimizedStringLengthV1)
}

func testMinimizedStringLength(t *testing.T, fn func(string) int) {
	testCases := []struct {
		name string
		s    string
		ans  int
	}{
		{"Preset Case 1", "aaabc", 3},
		{"Preset Case 2", "cbbd", 3},
		{"Preset Case 3", "baadccab", 4},
		{"Single character", "a", 1},
		{"All unique characters", "abcde", 5},
		{"Unique characters reversed", "edcba", 5},
		{"All same characters", "aaaaaa", 1},
		{"Adjacent repeating characters", "aaabbb", 2},
		{"Repeating characters in middle", "bcaaabc", 3},
		{"Repeating characters in end", "bcaaa", 3},
		{"Alternating characters", "ababab", 2},
		{"Duplicates at start and end", "abcda", 4},
		{"Multiple duplicates", "aabbccddeeff", 6},
		{"Mixed duplicates", "abacaba", 3},
		{"Maximum length string (100 chars)", repeatChar('a', 100), 1},
		{"Alternating 50 characters", repeatPattern("ab", 50), 2},
		{"Palindrome with duplicates", "racecar", 4},
		{"All characters same except one", "aaaaab", 2},
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
