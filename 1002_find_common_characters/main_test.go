package problem1002

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	testCases := []struct {
		words []string
		ans   [][]string
	}{
		{[]string{"bella", "label", "roller"}, [][]string{{"e", "l", "l"}, {"l", "e", "l"}, {"l", "l", "e"}}},
		{[]string{"cool", "lock", "cook"}, [][]string{{"c", "o"}, {"o", "c"}}},
		{[]string{"ccoool", "lock", "cccook"}, [][]string{{"c", "o"}, {"o", "c"}}},
		{[]string{"cool"}, [][]string{{"c", "l", "o", "o"}, {"l", "c", "o", "o"}, {"o", "l", "c", "o"}, {"o", "o", "l", "c"}, {"c", "o", "o", "l"}}},
		{[]string{"test", "liga", "moon"}, [][]string{}},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.words)
		t.Run(testName, func(t *testing.T) {
			ans := commonChars(testCase.words)
			if len(testCase.ans) == 0 {
				if len(ans) != 0 {
					t.Errorf("got %v, want %v", ans, testCase.ans)
				}
			} else {
				isFound := false
				for _, expectedAns := range testCase.ans {
					if reflect.DeepEqual(expectedAns, ans) {
						isFound = true
						break
					}
				}
				if !isFound {
					t.Errorf("got %v, want %v", ans, testCase.ans)
				}
			}
		})
	}
}
