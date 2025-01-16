package shufflearray_test

import (
	"reflect"
	"testing"

	shufflearray "shuffle_array"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{1, 2}, 1, []int{1, 2})
}

func checkFunc(t *testing.T, nums []int, index int, expected_result []int) {
	// GIVEN
	// WHEN
	result := shufflearray.Shuffle(nums, index)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
