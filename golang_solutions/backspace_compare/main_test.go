package backspacecompare_test

import (
	"reflect"
	"testing"

	backspacecompare "backspace_compare"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, "ab#c", "ad#c", true)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, "ab##", "c#d#", true)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, "a#c", "b", false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, "##", "###", true)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, "ab", "abc", false)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, "ab", "ab", true)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, "anb#", "and#", true)
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, "#anb#", "#and#", true)
}

func TestSuccess9(t *testing.T) {
	checkFunc(t, "f#nb", "u#nb", true)
}

func TestSuccess10(t *testing.T) {
	checkFunc(t, "", "", true)
}

func TestSuccess11(t *testing.T) {
	checkFunc(t, "vf#nb#t", "vy#nb#t", true)
}

func TestSuccess12(t *testing.T) {
	checkFunc(t, "vf#n#b#c", "vy#t#nert###", false)
}

func TestSuccess13(t *testing.T) {
	checkFunc(t, "vf###c", "cytn###", true)
}

func TestSuccess14(t *testing.T) {
	checkFunc(t, "bxj##tw", "bxo#j##tw", true)
}

func checkFunc(t *testing.T, s string, k string, expected_result bool) {
	result := backspacecompare.BackspaceCompare(s, k)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
