package movezeroes_test

import (
	"reflect"
	"testing"

	movezeroes "move_zeroes"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{0}, []int{0})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{}, []int{})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{1, 3, 12, 0, 0}, []int{1, 3, 12, 0, 0})
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{0, 1, 15, 0, 12, 3, 15}, []int{1, 15, 12, 3, 15, 0, 0})
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{0, 0, 0, -15, 10, -1, 0, 12, 0, 3, 15}, []int{-15, 10, -1, 12, 3, 15, 0, 0, 0, 0, 0})
}

func checkFunc(t *testing.T, nums []int, expected_result []int) {
	movezeroes.MoveZeroes(nums)
	if !reflect.DeepEqual(expected_result, nums) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, nums)
	}
}
