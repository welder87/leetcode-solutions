package sortpeople_test

import (
	"reflect"
	"testing"

	sortpeople "sort_people"
)

func TestSuccess1(t *testing.T) {
	checkFunc(t, []string{"Mary", "John", "Emma"}, []int{180, 165, 170}, []string{"Mary", "Emma", "John"})
}

func TestSuccess2(t *testing.T) {
	checkFunc(t, []string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}, []string{"Bob", "Alice", "Bob"})
}

func checkFunc(t *testing.T, names []string, heights []int, expected_result []string) {
	// GIVEN
	// WHEN
	result := sortpeople.SortPeople(names, heights)
	// THEN
	if !reflect.DeepEqual(expected_result, result) {
		t.Errorf("Expected element type: %v, got: %v", expected_result, result)
	}
}
