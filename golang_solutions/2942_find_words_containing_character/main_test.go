package problem2942

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	testCases := []struct {
		words []string
		x     rune
		ans   []int
	}{
		// preset cases
		{[]string{"leet", "code"}, 'e', []int{0, 1}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, 'a', []int{0, 2}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, 'z', []int{}},
		// common cases
		// corner cases
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.words, testCase.x)
		t.Run(testName, func(t *testing.T) {
			ans := findWordsContaining(testCase.words, byte(testCase.x))
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
