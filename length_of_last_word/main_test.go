package lengthoflastword_test

import (
	"testing"

	lengthoflastword "length_of_last_word"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "Hello World", 5)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "   fly me   to   the moon  ", 4)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "luffy is still joyboy", 6)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "", 0)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "  ", 0)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "  boy  ", 3)
}

func checkFunc(t *testing.T, s string, expected_result int) {
	result := lengthoflastword.LengthOfLastWord(s)
	if expected_result != result {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
