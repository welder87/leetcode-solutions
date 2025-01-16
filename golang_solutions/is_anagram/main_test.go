package isanagram_test

import (
	"reflect"
	"testing"

	isanagram "is_anagram"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "anagram", "nagaram", true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "rat", "car", false)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "anagram", "nagarama", false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "", "", true)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "a", "a", true)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "a", "b", false)
}

func checkFunc(t *testing.T, s1 string, s2 string, expected_result bool) {
	// GIVEN
	// WHEN
	result := isanagram.IsAnagram(s1, s2)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
