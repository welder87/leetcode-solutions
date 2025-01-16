package issubsequence_test

import (
	"reflect"
	"testing"

	issubsequence "is_subsequence"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "abc", "ahbgdc", true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "axc", "ahbgdc", false)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "ght", "nklmn", false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "abc", "ahgdcb", false)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "a", "ahgdcb", true)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "abc", "ahgdbc", true)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, "", "ahbgdc", true)
}

func checkFunc(t *testing.T, s string, ts string, expected_result bool) {
	result := issubsequence.IsSubsequence(s, ts)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
