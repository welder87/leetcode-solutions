package kidswithcandies_test

import (
	"reflect"
	"testing"

	kidswithcandies "kids_with_candies"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false})
}

func TestSuccess3(t *testing.T) {
	checkFunc(t, []int{12, 1, 12}, 10, []bool{true, false, true})
}

func TestSuccess4(t *testing.T) {
	checkFunc(t, []int{}, 10, []bool{})
}

func TestSuccess5(t *testing.T) {
	checkFunc(t, []int{9,9,9}, 1, []bool{true, true, true})
}



func checkFunc(t *testing.T, candies []int, extraCandies int, expected_result []bool) {
	// GIVEN
	// WHEN
	result := kidswithcandies.KidsWithCandies(candies, extraCandies)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf(
			"Input %v. Expected element type: %v, got: %v", candies, expected_result, result)
	}
}
