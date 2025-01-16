package sortsentence_test

import (
	"reflect"
	"testing"

	sortsentence "sort_sentence"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "is2 sentence4 This1 a3", "This is a sentence")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "Myself2 Me1 I4 and3", "Me Myself and I")
}

func checkFunc(t *testing.T, s string, expected_result string) {
	// GIVEN
	// WHEN
	result := sortsentence.SortSentence(s)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
