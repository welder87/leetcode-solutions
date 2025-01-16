package problem925

import (
	"fmt"
	"testing"
)

func TestIsLongPressedName(t *testing.T) {
	testCases := []struct {
		name  string
		typed string
		ans   bool
	}{
		{
			name:  "alex",
			typed: "aaleex",
			ans:   true,
		},
		{
			name:  "saeed",
			typed: "ssaaedd",
			ans:   true,
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.name, testCase.typed)
		t.Run(testName, func(t *testing.T) {
			ans := isLongPressedName(testCase.name, testCase.typed)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
