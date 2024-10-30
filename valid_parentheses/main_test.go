package validparentheses_test

import (
	"reflect"
	"testing"

	validparentheses "valid_parentheses"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "()", true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "()[]{}", true)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "(]", false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "([])", true)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "[({)]", false)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "[({}[()])]", true)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, "", true)
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, "[({}[(})])]", false)
}

func TestSuccess9(t *testing.T) {
	checkFunc(t, "])]", false)
}

func checkFunc(t *testing.T, s string, expected_result bool) {
	result := validparentheses.IsValid(s)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
