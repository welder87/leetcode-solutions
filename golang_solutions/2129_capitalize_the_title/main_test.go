package problem2129

import (
	"strings"
	"testing"
)

func TestCapitalizeTitle(t *testing.T) {
	testCapitalizeTitle(t, capitalizeTitle)
}

func TestCapitalizeTitleV1(t *testing.T) {
	testCapitalizeTitle(t, capitalizeTitleV1)
}

func testCapitalizeTitle(t *testing.T, fn func(string) string) {
	testCases := []struct {
		s    string
		ans  string
		name string
	}{
		{"capiTalIze tHe titLe", "Capitalize The Title", "Preset Case 1"},
		{"First leTTeR of EACH Word", "First Letter of Each Word", "Preset Case 2"},
		{"i lOve leetcode", "i Love Leetcode", "Preset Case 3"},
		{"ZW Cl pyR uoC", "zw cl Pyr Uoc", "Preset Case 4"},
		{"a", "a", "Single character word"},
		{"an", "an", "Single 2-character word"},
		{"the", "The", "Single 3-character word"},
		{"word", "Word", "4-character word should be capitalized"},
		{"A AN THE", "a an The", "All uppercase short words"},
		{
			"it is a test of and but for",
			"it is a Test of And But For",
			"Mixed case with boundary length",
		},
		{
			"the quick brown fox jumps over the lazy dog",
			"The Quick Brown Fox Jumps Over The Lazy Dog",
			"Long title with mixed lengths",
		},
		{
			"a " + strings.Repeat("verylongword", 10),
			"a " + "V" + strings.Repeat("verylongword", 10)[1:],
			"Very long word",
		},
		{"THE CAT AND THE HAT", "The Cat And The Hat", "All caps short words"},
		{"PROGRAMMING IS FUN", "Programming is Fun", "All caps long words"},
		{"an in on at to be", "an in on at to be", "Exactly 2 characters"},
		{"the and for but nor yet", "The And For But Nor Yet", "Exactly 3 characters"},
		{
			"this that what when where",
			"This That What When Where",
			"Exactly 4 characters",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.s); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
