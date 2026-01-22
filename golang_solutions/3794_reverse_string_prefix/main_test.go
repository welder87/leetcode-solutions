package problem3794

import (
	"testing"
)

func TestReversePrefix(t *testing.T) {
	testReversePrefix(t, reversePrefix)
}

func testReversePrefix(t *testing.T, fn func(string, int) string) {
	testCases := []struct {
		s    string
		k    int
		ans  string
		name string
	}{
		{"abcd", 2, "bacd", "Preset Case 1"},
		{"xyz", 3, "zyx", "Preset Case 2"},
		{"hey", 1, "hey", "Preset Case 2"},
		{"abcdef", 6, "fedcba", "Full reverse (even)"},
		{"abcde", 5, "edcba", "Full reverse (odd)"},
		{"abcdef", 3, "cbadef", "First three (even)"},
		{"abcde", 3, "cbade", "First three (odd)"},
		{"abcdef", 4, "dcbaef", "First four (even)"},
		{"abcde", 4, "dcbae", "First four (odd)"},
		{"a", 1, "a", "One char"},
		{"aa", 2, "aa", "Same chars"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s, testCase.k)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
