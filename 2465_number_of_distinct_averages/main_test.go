package problem2465

import (
	"fmt"
	"testing"
)

func TestDistinctAverages(t *testing.T) {
	testDistinctAverages(t, distinctAverages)
}

func testDistinctAverages(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{4, 1, 4, 0, 3, 5},
			ans:  2,
		},
		{
			nums: []int{1, 100},
			ans:  1,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func([]int) int
