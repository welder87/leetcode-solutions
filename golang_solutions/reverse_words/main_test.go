package reversewords_test

import (
	"testing"

	reversewords "reverse_words"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc")
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "Mr Ding", "rM gniD")
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "Mr", "rM")
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "", "")
}

func checkFunc(t *testing.T, s string, expected_result string) {
	result := reversewords.ReverseWords(s)
	if expected_result != result {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", s, expected_result, result)
	}
}
