package problem747

import (
	"fmt"
	"testing"
)

func TestDominantIndex(t *testing.T) {
	testDominantIndex(t, dominantIndex)
}

func testDominantIndex(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 6, 1, 0},
			ans:  1,
		},
		{
			nums: []int{1, 2, 3, 4},
			ans:  -1,
		},
		{
			nums: []int{1, 2},
			ans:  1,
		},
		{
			nums: []int{2, 3},
			ans:  -1,
		},
		{
			nums: []int{3, 3, 3, 3, 3, 4},
			ans:  -1,
		},
		{
			nums: []int{3, 3, 4, 3, 3},
			ans:  -1,
		},
		{
			nums: []int{3, 0, 3, 6, 3},
			ans:  3,
		},
		{
			nums: []int{0, 0, 0, 2, 0},
			ans:  3,
		},
		{
			nums: []int{0, 0, 0, 1},
			ans:  3,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) int
