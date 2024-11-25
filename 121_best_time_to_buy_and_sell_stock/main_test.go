package problem121

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		ans    int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			ans:    5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			ans:    0,
		},
		{
			prices: []int{7, 4, 5, 3, 1, 2},
			ans:    1,
		},
		{
			prices: []int{7, 1, 9, 3, 1, 10},
			ans:    9,
		},
		{
			prices: []int{1, 1, 1},
			ans:    0,
		},
		{
			prices: []int{1},
			ans:    0,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.prices)
		t.Run(testName, func(t *testing.T) {
			ans := maxProfit(testCase.prices)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
