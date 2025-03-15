package problem2951

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindPeaks(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  map[int]struct{}
	}{
		// preset cases
		{[]int{2, 4, 4}, map[int]struct{}{}},
		{[]int{1, 4, 3, 8, 5}, map[int]struct{}{1: {}, 3: {}}},
		// common cases
		{[]int{7, 6, 4, 5, 6, 7, 8}, map[int]struct{}{}},
		{[]int{7, 6, 5, 5, 8, 6, 3, 9, 8}, map[int]struct{}{4: {}, 7: {}}},
		{[]int{7, 6, 4, 3, 2, 1, 5}, map[int]struct{}{}},
		{[]int{1, 2, 3, 4, 4, 3, 2}, map[int]struct{}{}},
		{[]int{1, 2, 3, 4, 5, 4, 3, 2}, map[int]struct{}{4: {}}},
		// corner cases
		{[]int{7, 6, 4}, map[int]struct{}{}},
		{[]int{7, 8, 4}, map[int]struct{}{1: {}}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := findPeaks(testCase.nums)
			preparedAns := make(map[int]struct{}, len(ans))
			for _, val := range ans {
				preparedAns[val] = struct{}{}
			}
			if !reflect.DeepEqual(preparedAns, testCase.ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
