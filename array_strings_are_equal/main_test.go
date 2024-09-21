package arraystringsareequal_test

import (
	"reflect"
	"testing"

	arraystringsareequal "array_strings_are_equal"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []string{"ab", "c"}, []string{"a", "bc"}, true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []string{"a", "bc"}, []string{"", "abc"}, true)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []string{"a", "cb"}, []string{"ab", "c"}, false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []string{"abc", "d", "defg"}, []string{"abcddefg"}, true)
}

func checkFunc(t *testing.T, word1 []string, word2 []string, expected_result bool) {
	// GIVEN
	// WHEN
	result := arraystringsareequal.ArrayStringsAreEqual(word1, word2)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
