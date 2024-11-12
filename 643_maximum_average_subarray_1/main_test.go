package problem643

import (
	"fmt"
	"math"
	"testing"
)

const tolerance = 1e-5

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		ans  float64
	}{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			ans:  12.75000,
		},
		{
			nums: []int{5},
			k:    1,
			ans:  5.0,
		},
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    1,
			ans:  50.0,
		},
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    2,
			ans:  26.5,
		},
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    3,
			ans:  15.666666666666666,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := findMaxAverage(testCase.nums, testCase.k)
			if !floatEquals(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

func floatEquals(numA, numB float64) bool {
	return math.Abs(numA-numB) < tolerance
}
