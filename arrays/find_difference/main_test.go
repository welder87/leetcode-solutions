package finddifference_test

import (
	"reflect"
	"testing"

	finddifference "find_difference"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}})
}

func checkFunc(t *testing.T, nums1 []int, nums2 []int, expected_result [][]int) {
	result := finddifference.FindDifference(nums1, nums2)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
