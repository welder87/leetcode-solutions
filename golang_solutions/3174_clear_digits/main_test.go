package problem3174

import (
	"reflect"
	"testing"
)

func TestClearDigits(t *testing.T) {
	testClearDigits(t, clearDigits)
}

func testClearDigits(t *testing.T, fn func(string) string) {
	testCases := []struct {
		name string
		s    string
		ans  string
	}{
		{"Preset Case 1", "abc", "abc"},
		{"Preset Case 2", "cb34", ""},
		{"Single digit after letter", "a1", ""},
		{"Digit with preceding letter", "ab1", "a"},
		{"Multiple consecutive digits", "abc123", ""},
		{"Interleaved letters and digits", "a1b2c3", ""},
		{"Nested deletions 1", "ab12c", "c"},
		{"Nested deletions 2", "abc12d3", "a"},
		{"Digit at end with multiple letters", "hello5", "hell"},
		{"Multiple digits at end", "hello123", "he"},
		{"Cascading deletions 1", "ab1c2", "a"},
		{"Cascading deletions 2", "a1b2c3d4", ""},
		{"Complex cascading", "abc123def456", ""},
		{"Long string with pattern", "a1b1c1d1e1f1g1h1i1j1", ""},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
