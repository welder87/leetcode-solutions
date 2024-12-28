package problem1816

import (
	"fmt"
	"testing"
)

func TestTruncateSentence(t *testing.T) {
	testCases := []struct {
		s string
		k int
		ans   string
	}{
		{"Hello how are you Contestant", 4, "Hello how are you"},
		{"What is the solution to this problem", 4, "What is the solution"},
		{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.s, testCase.k)
		t.Run(testName, func(t *testing.T) {
			ans := truncateSentence(testCase.s, testCase.k)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
