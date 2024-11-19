package problem414

import (
	"fmt"
	"testing"
)

func TestThirdMax(t *testing.T) {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			nums: []int{3, 2, 1},
			ans:  1,
		},
		{
			nums: []int{1, 2},
			ans:  2,
		},
		{
			nums: []int{2, 2, 3, 1},
			ans:  1,
		},
		{
			nums: []int{2},
			ans:  2,
		},
		{
			nums: []int{1, 2, 3, 4, 1},
			ans:  2,
		},
		{
			nums: []int{10, 1, 3, 4, 1},
			ans:  3,
		},
		{
			nums: []int{2, 2, 2, 4, 1},
			ans:  1,
		},
		{
			nums: []int{1, 2, 2},
			ans:  2,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.nums)
		t.Run(testName, func(t *testing.T) {
			ans := thirdMax(testCase.nums)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
