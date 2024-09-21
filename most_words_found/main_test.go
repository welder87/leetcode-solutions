package mostwordsfound_test

import (
	"reflect"
	"testing"

	mostwordsfound "mostWordsFound"
)

func TestSuccess1(t *testing.T) {
	checkFunc(
		t,
		[]string{
			"alice and bob love leetcode",
			"i think so too",
			"this is great thanks very much",
		},
		6,
	)
}


func TestSuccess2(t *testing.T) {
	checkFunc(
		t,
		[]string{"please wait", "continue to fight", "continue to win"},
		3,
	)
}

func checkFunc(t *testing.T, sentences []string, expected_result int) {
	// GIVEN
	// WHEN
	result := mostwordsfound.MostWordsFound(sentences)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", sentences, expected_result, result)
	}
}
