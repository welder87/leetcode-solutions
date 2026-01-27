package problem3110

import (
	"reflect"
	"testing"
)

func TestScoreOfString(t *testing.T) {
	testScoreOfString(t, scoreOfString)
}

func TestScoreOfStringV1(t *testing.T) {
	testScoreOfString(t, scoreOfStringV1)
}

func testScoreOfString(t *testing.T, fn func(string) int) {
	testCases := []struct {
		s    string
		ans  int
		name string
	}{
		{"aaz", 25, ""},
		{"hello", 13, "Preset Case 1"},
		{"zaz", 50, "Preset Case 2"},
		{"aa", 0, "Two identical characters"},
		{"ab", 1, "Two consecutive alphabet characters"},
		{"ab", 1, "Reverse consecutive"},
		{"aaaaa", 0, "All same character"},
		{"zzzzzzzz", 0, "All same character long"},
		{"abcde", 4, "Increasing sequence"},
		{"abcdefghijklmnopqrstuvwxyz", 25, "Full alphabet increasing"},
		{"edcba", 4, "Decreasing sequence"},
		{"zyxwvutsrqponmlkjihgfedcba", 25, "Full alphabet decreasing"},
		{"azaza", 100, "Alternating high-low"},
		{"zazaz", 100, "Alternating low-high"},
		{"az", 25, "Maximum difference"},
		{"za", 25, "Maximum to minimum"},
		{"ababababab", 9, "Long repeating pattern"},
		{"abcba", 4, "Palindrome odd"},
		{"abccba", 4, "Palindrome even"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
