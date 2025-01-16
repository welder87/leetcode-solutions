package problem1833

import (
	"fmt"
	"testing"
)

func TestMaxIceCream(t *testing.T) {
	testCases := []struct {
		costs []int
		coins int
		ans   int
	}{
		{
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			ans:   4,
		},
		{
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			ans:   0,
		},
		{
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			ans:   6,
		},
		{
			costs: []int{1, 3, 2},
			coins: 7,
			ans:   3,
		},
		{
			costs: []int{3, 1, 3},
			coins: 7,
			ans:   3,
		},
		{
			costs: []int{3, 1, 3},
			coins: 7,
			ans:   3,
		},
		{
			costs: []int{3, 1, 3, 3},
			coins: 7,
			ans:   3,
		},
		{
			costs: []int{3, 2, 3, 3},
			coins: 7,
			ans:   2,
		},
		{
			costs: []int{3},
			coins: 7,
			ans:   1,
		},
		{
			costs: []int{7},
			coins: 7,
			ans:   1,
		},
		{
			costs: []int{8},
			coins: 7,
			ans:   0,
		},
		{
			costs: []int{8, 9, 8, 12},
			coins: 7,
			ans:   0,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.costs, testCase.coins)
		t.Run(testName, func(t *testing.T) {
			ans := maxIceCream(testCase.costs, testCase.coins)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
