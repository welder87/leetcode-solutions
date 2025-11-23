package problem283

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testMoveZeroes(t, moveZeroes)
}

func TestMoveZeroesV1(t *testing.T) {
	testMoveZeroes(t, moveZeroesV1)
}

func testMoveZeroes(t *testing.T, fn func([]int)) {
	testCases := []struct {
		testName string
		nums     []int
		ans      []int
	}{
		// PRESET CASES
		{"Preset_case_1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"Preset_case_2", []int{0}, []int{0}},
		// COMMON CASES
		{"Trailing_zeros", []int{1, 3, 12, 0, 0}, []int{1, 3, 12, 0, 0}},
		{"Leading_zeros", []int{0, 0, 0, 1, 2, 3}, []int{1, 2, 3, 0, 0, 0}},
		{"Embedded_zeros", []int{1, 0, 2, 0, 3, 0}, []int{1, 2, 3, 0, 0, 0}},
		{
			"Alternating_zeros_and_non_zeros_(zero_start)",
			[]int{0, 1, 0, 2, 0, 3, 0, 4},
			[]int{1, 2, 3, 4, 0, 0, 0, 0},
		},
		{
			"Alternating_zeros and non_zeros (non_zero_start)",
			[]int{1, 0, 2, 0, 3, 0, 4},
			[]int{1, 2, 3, 4, 0, 0, 0},
		},
		{
			"Leading_and_trailing_zeros",
			[]int{0, 0, 0, 1, 3, 12, 0, 0},
			[]int{1, 3, 12, 0, 0, 0, 0, 0},
		},
		{
			"Large_numbers",
			[]int{1000000, 0, 2147483648, 0, 0, 42},
			[]int{1000000, 2147483648, 42, 0, 0, 0},
		},
		// CORNER CASES
		{"No_zeros", []int{1, 6, 1, 24}, []int{1, 6, 1, 24}},
		{"All_zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"Single_Zero", []int{0}, []int{0}},
		{"Single_Non_Zero", []int{5}, []int{5}},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			fn(tc.nums)
			if !reflect.DeepEqual(tc.nums, tc.ans) {
				t.Errorf("got %v, want %v", tc.nums, tc.ans)
			}
		})
	}
}
