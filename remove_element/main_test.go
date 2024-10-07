package removeelement_test

import (
	"reflect"
	"testing"

	removeelement "remove_element"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{3, 2, 2, 3}, 3, 2)
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5)
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{2, 1, 2, 2, 3, 0, 4, 2}, 2, 4)
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{}, 2, 0)
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{1}, 2, 1)
}

func TestSuccess6(t *testing.T) {
	checkFunc(t, []int{2}, 2, 0)
}

func TestSuccess7(t *testing.T) {
	checkFunc(t, []int{2, 1, 0, 2, 3, 0, 4, 2}, 0, 6)
}

func TestSuccess8(t *testing.T) {
	checkFunc(t, []int{2, 2, 2}, 2, 0)
}

func TestSuccess9(t *testing.T) {
	checkFunc(t, []int{0, 1, 0, 2, 3, 0, 4, 1}, 2, 7)
}

func TestSuccess10(t *testing.T) {
	checkFunc(t, []int{0, 1, 0, 2, 3, 0, 4, 1}, 5, 8)
}

func checkFunc(t *testing.T, nums []int, val int, expected_result int) {
	result := removeelement.RemoveElement(nums, val)
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
