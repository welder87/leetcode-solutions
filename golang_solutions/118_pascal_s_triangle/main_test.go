package problem118

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		numRows int
		ans     [][]int
	}{
		{
			numRows: 5,
			ans:     [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows: 1,
			ans:     [][]int{{1}},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.numRows)
		t.Run(testName, func(t *testing.T) {
			ans := generate(testCase.numRows)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
