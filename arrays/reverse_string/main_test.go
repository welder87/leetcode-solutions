package reversestring_test

import (
	"reflect"
	"testing"

	reversestring "reverse_string"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []byte("hello"), []byte("olleh"))
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []byte("Hannah"), []byte("hannaH"))
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []byte("Chain"), []byte("niahC"))
}

func checkFunc(t *testing.T, s []byte, expected_result []byte) {
	reversestring.ReverseString(s)
	if !reflect.DeepEqual(expected_result, s) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", s, expected_result, s)
	}
}
