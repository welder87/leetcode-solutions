package sortarraybyparity_test

import (
	"reflect"
	"testing"

	sortarraybyparity "sort_array_by_parity"
)

func TestSuccess1(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{4, 2, 3, 1})
	expected_result = append(expected_result, []int{2, 4, 1, 3})
	expected_result = append(expected_result, []int{4, 2, 1, 3})
	checkFunc(t, []int{3, 1, 2, 4}, expected_result)
}

func TestSuccess2(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{0})
	checkFunc(t, []int{0}, expected_result)
}

func TestSuccess3(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{2, 4, 8})
	expected_result = append(expected_result, []int{2, 8, 4})
	expected_result = append(expected_result, []int{4, 8, 2})
	checkFunc(t, []int{2, 4, 8}, expected_result)
}

func TestSuccess4(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{3, 1, 5})
	expected_result = append(expected_result, []int{5, 3, 1})
	expected_result = append(expected_result, []int{1, 3, 5})
	checkFunc(t, []int{3, 1, 5}, expected_result)
}

func TestSuccess5(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{0, 4, 1})
	expected_result = append(expected_result, []int{4, 0, 1})
	checkFunc(t, []int{0, 1, 4}, expected_result)
}

func TestSuccess6(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{4, 2, 3, 1})
	expected_result = append(expected_result, []int{2, 4, 1, 3})
	expected_result = append(expected_result, []int{4, 2, 1, 3})
	expected_result = append(expected_result, []int{2, 4, 3, 1})
	checkFunc(t, []int{2, 4, 3, 1}, expected_result)
}

func TestSuccess7(t *testing.T) {
	expected_result := [][]int{}
	expected_result = append(expected_result, []int{4, 2, 3})
	expected_result = append(expected_result, []int{2, 4, 3})
	checkFunc(t, []int{3, 2, 4}, expected_result)
}

func checkFunc(t *testing.T, nums []int, expected_result [][]int) {
	result := sortarraybyparity.SortArrayByParity(nums)
	check := false
	for _, subSlice := range expected_result {
		if reflect.DeepEqual(subSlice, result) {
			check = true
			break
		}
	}
	if !check {
		t.Errorf("Error")
	}
}
