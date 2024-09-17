package reversevowels_test

import (
	"testing"

	reversevowels "reverse_vowels"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "IceCreAm", "AceCreIm")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "leetcode", "leotcede")
}

func checkFunc(t *testing.T, s string, expected_result string) {
	result := reversevowels.ReverseVowels(s)
	if expected_result == s {
		t.Errorf("Input %v. Expected element type: %v, got: %v", s, expected_result, result)
	}
}
