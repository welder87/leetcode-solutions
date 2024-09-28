package reverseonlyletters_test

import (
	"reflect"
	"testing"

	reverseonlyletters "reverse_only_letters"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "ab-cd", "dc-ba")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "a-bC-dEf-ghIj", "j-Ih-gfE-dCba")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!")
}

func checkFunc(t *testing.T, s string, expected_result string) {
	result := reverseonlyletters.ReverseOnlyLetters(s)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Input %v. Expected element type: %v, got: %v", s, expected_result, result)
	}
}
