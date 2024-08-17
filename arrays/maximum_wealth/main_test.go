package maximumwealth_test

import (
	"reflect"
	"testing"

	maximumwealth "maximum_wealth"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, [][]int{{1, 2, 3}, {3, 1, 1}}, 6)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, [][]int{{1, 5}, {7, 3}, {3, 5}}, 10)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17)
}

func checkFunc(t *testing.T, nums [][]int, expected_result int) {
	// GIVEN
	// WHEN
	result := maximumwealth.MaximumWealth(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
