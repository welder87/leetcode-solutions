package sortedsquares_test

import (
	"reflect"
	"testing"

	sortedsquares "sorted_squares"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{-4, 0, 3, 10}, []int{0, 9, 16, 100})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{0, 3, 10, 15}, []int{0, 9, 100, 225})
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{-15, -7, -3, -1}, []int{1, 9, 49, 225})
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{}, []int{})
}

func checkFunc(t *testing.T, nums []int, expected_result []int) {
	// GIVEN
	// WHEN
	result := sortedsquares.SortedSquares(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
