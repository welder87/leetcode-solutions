package removeduplicates_test

import (
	"reflect"
	"testing"

	removeduplicates "remove_duplicates"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 1, 2}, 2)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5)
}

func checkFunc(t *testing.T, nums []int, expected_result int) {
	result := removeduplicates.RemoveDuplicates(nums)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
