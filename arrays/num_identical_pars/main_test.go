package numidenticalpars_test

import (
	"reflect"
	"testing"

	numidenticalpars "num_identical_pars"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 1, 1, 3}, 4)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{}, 0)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{1}, 0)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{1, 1}, 1)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{1, -1, 1}, 1)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{1, 1, 1, 1}, 6)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, []int{1, 2, 3}, 0)
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, []int{15, -2, 0, 3, -17, -2, 0, 86, 0, -2}, 6)
}

func checkFunc(t *testing.T, nums []int, expected_result int) {
	// GIVEN
	// WHEN
	result := numidenticalpars.NumIdenticalPairs(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
