package minimumabsdifference_test

import (
	"reflect"
	"testing"

	minimumabsdifference "minimum_abs_difference"
)

func TestSuccess1(t *testing.T) {
	var expected_result [][]int
	expected_result = append(expected_result, []int{1, 2})
	expected_result = append(expected_result, []int{2, 3})
	expected_result = append(expected_result, []int{3, 4})
	checkFunc(t, []int{4, 2, 1, 3}, expected_result)
}

func TestSuccess2(t *testing.T) {
	var expected_result [][]int
	expected_result = append(expected_result, []int{-14, -10})
	expected_result = append(expected_result, []int{19, 23})
	expected_result = append(expected_result, []int{23, 27})
	checkFunc(t, []int{3, 8, -10, 23, 19, -4, -14, 27}, expected_result)
}

func TestSuccess3(t *testing.T) {
	var expected_result [][]int
	expected_result = append(expected_result, []int{1, 3})
	checkFunc(t, []int{1, 3, 6, 10, 15}, expected_result)
}

func checkFunc(t *testing.T, arr []int, expected_result [][]int) {
	result := minimumabsdifference.MinimumAbsDifference(arr)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
