package problem1920

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := buildArray(testCase.nums)
			if !reflect.DeepEqual(testCase.expected, ans) {
				t.Errorf("got %v, want %v", ans, testCase.expected)
			}
		})
	}
}
