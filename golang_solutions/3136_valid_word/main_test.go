package problem3136

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	testIsValid(t, isValid)
}

func TestIsValidV1(t *testing.T) {
	testIsValid(t, isValidV1)
}

func testIsValid(t *testing.T, fn func(string) bool) {
	testCases := []struct {
		s    string
		ans  bool
		name string
	}{
		{"234Adas", true, "Preset Case 1"},
		{"b3", false, "Preset Case 2"},
		{"a3$e", false, "Preset Case 3"},
		{"aya", true, "Preset Case 4"},
		{"IS", false, "Preset Case 5"},
		{"a1b", true, "Exactly 3 characters (vowel, consonant, digit)"},
		{"a1e", false, "Exactly 3 characters (only vowels and digit)"},
		{"b2c", false, "Exactly 3 characters (only consonants and digits)"},
		{"ab", false, "Length 2 - too short"},
		{"a", false, "Length 1 - too short"},
		{
			"abcdefghij1234567890",
			true,
			"Maximum length (20 chars) (has vowels and consonants)",
		},
		{"123456", false, "Only digits - no vowels or consonants"},
		{"123@456", false, "Digits with special char"},
		{"@#$", false, "Special character only"},
		{"abc@123", false, "Mixed valid and invalid chars"},
		{"BCDFG", false, "All valid chars (no vowel)"},
		{"aeiou", false, "All valid chars (no consonant)"},
		{"a123bc", true, "Single lowercase vowel with consonants"},
		{"E123bc", true, "Single uppercase vowel with consonants"},
		{"aeiouAEIOU", false, "All vowels - no consonants"},
		{"aei123ou", false, "Mixed vowels and digits - no consonants"},
		{"apple123", true, "Vowel at beginning"},
		{"123bana", true, "Vowel at end"},
		{"aeioub", true, "All vowels except one consonant"},
		{"b123ae", true, "Single consonant with vowels"},
		{"bcdfghjklmnpqrstvwxyz", false, "All consonants - no vowels"},
		{"bcd123fgh", false, "Consonants with digits - no vowels"},
		{"ball123", true, "Consonant at beginning"},
		{"123aeiob", true, "Consonant at end"},
		{"bcdfghja", true, "All consonants except one vowel"},
		{"HeLlO123", true, "Mixed case - valid"},
		{"AEIOUBCD123", true, "Uppercase vowels"},
		{"BCDFG", false, "Uppercase consonants only"},
		{"AEIOU", false, "Uppercase vowels only"}, // no consonant
		{"aeiou", false, "Lowercase vowels only"},
		{"abc123def", true, "Mixed alphanumeric with both requirements"},
		{"supercalifragilisticexpialidocious123", true, "Long valid word"},
		{"a1b2c3d4e5f6", true, "Numbers interspersed"},
		{"123hello", true, "Starts with number"},
		{"hello123", true, "Ends with number"},
		{"bcdfghjklmnpa", true, "Only one vowel among many consonants"},
		{"aeiouaeioub", true, "Only one consonant among many vowels"},
		{"bcd123fgh", false, "No vowels with digits"},
		{"aei123ou", false, "No consonants with digits"},
		{"abc@123def", false, "Special character in middle"},
		{"abc@123", false, "Single allowed special char '@' - invalid"},
		{"abc#123", false, "Single allowed special char '#' - invalid"},
		{"abc$123", false, "Single allowed special char '$' - invalid"},
		{"123", false, "Minimum length with digits only"},
		{"bcd", false, "Minimum length with letters only (no vowel)"},
		{"aei", false, "Minimum length with letters only (no consonant)"},
		{"ab1", true, "Minimum length valid"},
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
