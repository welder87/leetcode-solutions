package problem704

import (
	"strings"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	testToLowerCase(t, toLowerCase)
}

func TestToLowerCaseV1(t *testing.T) {
	testToLowerCase(t, toLowerCaseV1)
}

func testToLowerCase(t *testing.T, fn func(string) string) {
	testCases := []struct {
		name string
		s    string
		ans  string
	}{
		{"Preset case 1", "Hello", "hello"},
		{"Preset case 2", "here", "here"},
		{"Preset case 3", "LOVELY", "lovely"},
		{"Mixed case letters", "Hello World", "hello world"},
		{"Already lowercase", "hello", "hello"},
		{"Single uppercase letter", "A", "a"},
		{"Single lowercase letter", "z", "z"},
		{"With numbers and special characters", "HELLO123!@#", "hello123!@#"},
		{"Empty string", "", ""},
		{"Spaces only", "   ", "   "},
		{"Tabs and newlines", "\tHELLO\nWORLD\t", "\thello\nworld\t"},
		{
			"Complete uppercase alphabet",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"abcdefghijklmnopqrstuvwxyz",
		},
		{"Long uppercase string",
			strings.Repeat("TEST ", 100),
			strings.Repeat("test ", 100),
		},
		{
			"Case variations",
			"AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz",
			"aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz",
		},
		{"String with only special characters", "!@#$%^&*()_+", "!@#$%^&*()_+"},
		{"String starting with uppercase", "TestString", "teststring"},
		{"String ending with uppercase", "testSTRING", "teststring"},
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
