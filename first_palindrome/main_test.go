package firstpalindrome_test

import (
	"reflect"
	"testing"

	firstpalindrome "first_palindrome"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []string{"abc", "car", "ada", "racecar", "cool"}, "ada")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []string{"notapalindrome", "racecar"}, "racecar")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []string{"def", "ghi"}, "")
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []string{"adttda", "ghi"}, "adttda")
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []string{"ghi", "adtutda"}, "adtutda")
}

func checkFunc(t *testing.T, words []string, expected_result string) {
	result := firstpalindrome.FirstPalindrome(words)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
