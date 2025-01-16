package problem1652

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		ans  []int
	}{
		{
			nums: []int{5, 7, 1, 4},
			k:    3,
			ans:  []int{12, 10, 16, 13},
		},
		{
			nums: []int{1, 2, 3, 4},
			k:    0,
			ans:  []int{12, 10, 16, 13},
		},
		{
			nums: []int{2, 4, 9, 3},
			k:    -2,
			ans:  []int{12, 10, 16, 13},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.nums, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := decrypt(testCase.nums, testCase.k)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
