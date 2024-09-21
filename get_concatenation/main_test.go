package getconcatenation_test

import (
	"reflect"
	"testing"

	getconcatenation "get_concatenation"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{1}, []int{1, 1})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{}, []int{})
}

func checkFunc(t *testing.T, nums []int, expected_result []int) {
	// GIVEN
	// WHEN
	result := getconcatenation.GetConcatenation(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
