package minimumaverage_test

import (
	"reflect"
	"testing"

	minimumaverage "minimum_average"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{7, 8, 3, 4, 15, 13, 4, 1}, 5.5)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1, 9, 8, 3, 10, 5}, 5.5)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 7, 8, 9}, 5.0)
}

func checkFunc(t *testing.T, nums []int, expected_result float64) {
	result := minimumaverage.MinimumAverage(nums)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
