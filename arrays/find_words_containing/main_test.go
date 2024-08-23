package findwordscontaining_test

import (
	"reflect"
	"testing"

	findwordscontaining "find_words_containing"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []string{"leet", "code"}, []byte("e")[0], []int{0, 1})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []string{"abc", "bcd", "aaaa", "cbc"}, []byte("a")[0], []int{0, 2})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []string{"abc","bcd","aaaa","cbc"}, []byte("z")[0], []int{})
}

func checkFunc(t *testing.T, words []string, x byte, expected_result []int) {
	// GIVEN
	// WHEN
	result := findwordscontaining.FindWordsContaining(words, x)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
