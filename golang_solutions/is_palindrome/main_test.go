package ispalindrome_test

import (
	"reflect"
	"testing"

	ispalindrome "is_palindrome"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "A man, a plan, a canal: Panama", true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "race a car", false)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, " ", true)
}

func checkFunc(t *testing.T, s string, expected_result bool) {
	result := ispalindrome.IsPalindrome(s)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
