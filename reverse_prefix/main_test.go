package reverseprefix_test

import (
	"reflect"
	"testing"

	reverseprefix "reverse_prefix"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "abcdefd", []byte("d")[0], "dcbaefd")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "xyxzxe", []byte("z")[0], "zxyxxe")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "abcd", []byte("z")[0], "abcd")
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "abcd", []byte("d")[0], "dcba")
}

func checkFunc(t *testing.T, word string, ch byte, expected_result string) {
	result := reverseprefix.ReversePrefix(word, ch)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Input %v. Expected element type: %v, got: %v", word, expected_result, result)
	}
}
