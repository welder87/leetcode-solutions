package mergealternately_test

import (
	"reflect"
	"testing"

	mergealternately "merge_alternately"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "abc", "pqr", "apbqcr")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "ab", "pqrs", "apbqrs")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "abcd", "pq", "apbqcd")
}

func checkFunc(t *testing.T, word1 string, word2 string, expected_result string) {
	result := mergealternately.MergeAlternately(word1, word2)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
