package strstr_test

import (
	"reflect"
	"testing"

	strstr "str_str"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "sadbutsad", "sad", 0)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "leetcode", "leeto", -1)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "liloslilas", "lila", 5)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "aaa", "aaaa", -1)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "mississippi", "issipi", -1)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "hello", "ll", 2)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, "ll", "ll", 0)
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, "lo", "ll", -1)
}

func TestSuccess9(t *testing.T) {
	checkFunc(t, "abc", "c", 2)
}

func TestSuccess10(t *testing.T) {
	checkFunc(t, "lililoslililas", "lilila", 7)
}

func checkFunc(t *testing.T, haystack string, needle string, expected_result int) {
	result := strstr.StrStr(haystack, needle)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
