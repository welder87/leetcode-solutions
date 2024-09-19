package intersectiontwoarrays2_test

import (
	"reflect"
	"testing"

	intersectiontwoarrays2 "intersection_two_arrays2"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1, 2, 2, 1}, []int{2, 2}, [][]int{{2, 2}})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, [][]int{{9, 4}, {4, 9}})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{1, 2, 3}, []int{4, 6, 5}, [][]int{{}})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{}, []int{}, [][]int{{}})
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{2, 2, 2, 2}, []int{2, 2}, [][]int{{2, 2}})
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{1, 2, 1, 2}, []int{2, 2}, [][]int{{2, 2}})
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, []int{1, 2, 1, 2, 5, 2, 1}, []int{2, 17, 2, 5}, [][]int{{2, 2, 5}, {2, 5, 2}, {5, 2, 2}})
}

func checkFunc(t *testing.T, nums1 []int, nums2 []int, expected_result [][]int) {
	result := intersectiontwoarrays2.Intersect(nums1, nums2)
	isFound := false
	for _, expected := range expected_result {
		if isFound = reflect.DeepEqual(expected, result); isFound {
			break
		}
	}
	if !isFound {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
