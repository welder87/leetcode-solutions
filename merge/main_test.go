package merge_test

import (
	"reflect"
	"testing"

	"merge"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1}, 1, []int{}, 0, []int{1})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{0}, 0, []int{1}, 1, []int{1})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 3, 0, 0, 0}, 4, []int{1, 2, 3}, 3, []int{1, 1, 2, 2, 3, 3, 3})
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{4, 5, 5, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 5})
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6})
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, []int{1, 2, 4, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 4, 4, 5, 6})
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, []int{1, 2, 2, 0, 0, 0}, 3, []int{1, 5, 6}, 3, []int{1, 1, 2, 2, 5, 6})
}

func checkFunc(t *testing.T, nums1 []int, m int, nums2 []int, n int, expected_result []int) {
	merge.Merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(expected_result, nums1) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, nums1)
	}
}
