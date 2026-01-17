package problem520

import (
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {
	testDetectCapitalUse(t, detectCapitalUse)
}

func testDetectCapitalUse(t *testing.T, fn func(string) bool) {
	testCases := []struct {
		name string
		s    string
		ans  bool
	}{
		{"Preset case 1", "USA", true},
		{"Preset case 2", "FlaG", false},
		{"Preset case 3", "leetcode", true},
		{"Preset case 4", "Google", true},
		{"Single uppercase letter", "A", true},
		{"Single lowercase letter", "a", true},
		{"Two letters first capital", "Ab", true},
		{"Two letters all uppercase", "AB", true},
		{"Two letters all lowercase", "ab", true},
		{"Lowercase first then uppercase", "mL", false},
		{"Capital first then lowercase then capital", "GooGle", false},
		{"Multiple capitals in middle", "TesTing", false},
		{"Only second letter capital", "aBc", false},
		{"Capital at end", "abcD", false},
		{"Two letters wrong case", "aB", false},
		{"All same letter uppercase", "AAAAAAAAAAAAA", true},
		{"All same letter lowercase", "bbbbbbbbbbbbbbb", true},
		{"Long valid uppercase word", "INTERNATIONAL", true},
		{"Long valid lowercase word", "implementation", true},
		{"Long valid title case word", "Capitalization", true},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
