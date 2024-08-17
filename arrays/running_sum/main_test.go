package runningsum_test

import (
	"reflect"
	"testing"

	runningsum "running_sum"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{1,2,3,4}, []int{1,3,6,10})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{1,1,1,1,1}, []int{1,2,3,4,5})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{3,1,2,10,1}, []int{3,4,6,16,17})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{-1,0,-5, 20}, []int{-1,-1,-6, 14})
}


func checkFunc(t *testing.T, nums []int, expected_result []int) {
	// GIVEN
	// WHEN
	result := runningsum.RunningSum(nums)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", nums, expected_result, result)
	}
}
