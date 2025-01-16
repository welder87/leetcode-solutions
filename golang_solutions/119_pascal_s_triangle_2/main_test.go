package problem119

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	testCases := []struct {
		rowIndex int
		ans      []int
	}{
		{
			rowIndex: 3,
			ans:      []int{1, 3, 3, 1},
		},
		{
			rowIndex: 0,
			ans:      []int{1},
		},
		{
			rowIndex: 1,
			ans:      []int{1, 1},
		},
		{
			rowIndex: 4,
			ans:      []int{1, 4, 6, 4, 1},
		},
		{
			rowIndex: 5,
			ans:      []int{1, 5, 10, 10, 5, 1},
		},
		{
			rowIndex: 6,
			ans:      []int{1, 6, 15, 20, 15, 6, 1},
		},
		{
			rowIndex: 7,
			ans:      []int{1, 7, 21, 35, 35, 21, 7, 1},
		},
		{
			rowIndex: 12,
			ans:      []int{1, 12, 66, 220, 495, 792, 924, 792, 495, 220, 66, 12, 1},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.rowIndex)
		t.Run(testName, func(t *testing.T) {
			ans := getRow(testCase.rowIndex)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
