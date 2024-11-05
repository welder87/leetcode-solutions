package problem1961

import (
	"fmt"
	"testing"
)

func TestIsPrefixString(t *testing.T) {
	testCases := []struct {
		s     string
		words []string
		want  bool
	}{
		{"iloveleetcode", []string{"i", "love", "leetcode", "apples"}, true},
		{"iloveleetcode", []string{"apples", "i", "love", "leetcode"}, false},
		{"iloveleetcode", []string{"i", "love", "apples", "leetcode"}, false},
		{"goodday", []string{"good"}, false},
		{"a", []string{"aa", "aaaa", "banana"}, false},
		{"aaa", []string{"aa", "aaa", "fjaklfj"}, false},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.s, testCase.words)
		t.Run(testName, func(t *testing.T) {
			ans := IsPrefixString(testCase.s, testCase.words)
			if ans != testCase.want {
				t.Errorf("got %v, want %v", ans, testCase.want)
			}
		})
	}
}
