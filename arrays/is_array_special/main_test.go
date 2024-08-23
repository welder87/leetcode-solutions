package isarrayspecial_test

import (
	"reflect"
	"testing"

	isarrayspecial "is_array_special"
)



func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1}, true)
}
func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{2,1,4}, true)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{4,3,1,6}, false)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{}, true)
}

func checkFunc(t *testing.T, nums []int, expected_result bool) {
	// GIVEN
	// WHEN
	result := isarrayspecial.IsArraySpecial(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
