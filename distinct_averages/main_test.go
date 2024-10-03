package distinctaverages_test

import (
	"reflect"
	"testing"

	distinctaverages "distinct_averages"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{4, 1, 4, 0, 3, 5}, 2)
}

func checkFunc(t *testing.T, nums []int, expected_result int) {
	result := distinctaverages.DistinctAverages(nums)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
