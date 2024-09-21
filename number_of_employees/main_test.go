package numberofemployees_test

import (
	"reflect"
	"testing"

	numberofemployees "number_of_employees_who_met_target"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{0, 1, 2, 3, 4}, 2, 3)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{5, 1, 4, 2, 2}, 6, 0)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{0, 0, 0, 0, 0}, 2, 0)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{2, 2, 2}, 2, 3)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{3, 3, 3}, 2, 3)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{}, 2, 0)
}

func checkFunc(t *testing.T, nums []int, target int, expected_result int) {
	// GIVEN
	// WHEN
	result := numberofemployees.NumberOfEmployeesWhoMetTarget(nums, target)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
