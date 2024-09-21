package truncatesentence_test

import (
	"reflect"
	"testing"

	truncatesentence "truncate_sentence"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "Hello how are you Contestant", 4, "Hello how are you")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "What is the solution to this problem", 4, "What is the solution")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "chopper is not a tanuki", 5, "chopper is not a tanuki")
}

func checkFunc(t *testing.T, s string, k int, expected_result string) {
	// GIVEN
	// WHEN
	result := truncatesentence.TruncateSentence(s, k)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", s, expected_result, result)
	}
}
