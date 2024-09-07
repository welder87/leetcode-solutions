package findthedifference_test

import (
	"reflect"
	"testing"

	findthedifference "find_the_difference"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "abcd", "abcde", []byte("e")[0])
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "", "y", []byte("y")[0])
}

func checkFunc(t *testing.T, s1 string, s2 string, expected_result byte) {
	result := findthedifference.FindTheDifference(s1, s2)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
