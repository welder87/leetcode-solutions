package problem2733

import (
	"fmt"
	"testing"
)

type fn func([]int) int

func TestFindNonMinOrMax(t *testing.T) {
	testFindNonMinOrMax(t, findNonMinOrMax)
}

func TestFindNonMinOrMaxV1(t *testing.T) {
	testFindNonMinOrMax(t, findNonMinOrMaxV1)
}

func TestFindNonMinOrMaxV2(t *testing.T) {
	testFindNonMinOrMax(t, findNonMinOrMaxV2)
}

func testFindNonMinOrMax(t *testing.T, function fn) {
	testCases := []struct {
		nums []int
		ans  map[int]struct{}
	}{
		// preset cases
		{[]int{3, 2, 1, 4}, map[int]struct{}{2: {}, 3: {}}},
		{[]int{1, 2}, map[int]struct{}{-1: {}}},
		{[]int{2, 1, 3}, map[int]struct{}{2: {}}},
		{[]int{2, 4, 25}, map[int]struct{}{4: {}}},
		// common cases
		{[]int{3, 50, 1, 100, 25}, map[int]struct{}{50: {}, 25: {}, 3: {}}},
		// corner cases
		{[]int{1}, map[int]struct{}{-1: {}}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.nums)
			if _, ok := testCase.ans[ans]; !ok {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
