package countpairs_test

import (
	countpairs "count_pairs"
	"reflect"
	"testing"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{-6, 2, 5, -2, -7, -1, 3}, -2, 10)
}

func checkFunc(t *testing.T, nums []int, target int, expected_result int) {
	// GIVEN
	// WHEN
	result := countpairs.CountPairs(nums, target)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
