package problem2485

import (
	"fmt"
	"testing"
)

func TestPivotInteger(t *testing.T) {
	testPivotInteger(t, pivotInteger)
}

func TestPivotIntegerV1(t *testing.T) {
	testPivotInteger(t, pivotIntegerV1)
}

func testPivotInteger(t *testing.T, function fn) {
	testCases := []struct {
		n   int
		ans int
	}{
		{n: 8, ans: 6},
		{n: 1, ans: 1},
		{n: 4, ans: -1},
		{n: 2, ans: -1},
		{n: 3, ans: -1},
		{n: 5, ans: -1},
		{n: 6, ans: -1},
		{n: 7, ans: -1},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.n)
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.n)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func(int) int
