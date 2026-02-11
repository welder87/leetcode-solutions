package problem1880

import (
	"reflect"
	"strings"
	"testing"
)

func TestIsSumEqual(t *testing.T) {
	testIsSumEqual(t, isSumEqual)
}

func TestIsSumEqualV1(t *testing.T) {
	testIsSumEqual(t, isSumEqualV1)
}

func testIsSumEqual(t *testing.T, fn func(string, string, string) bool) {
	testCases := []struct {
		firstWord  string
		secondWord string
		targetWord string
		ans        bool
		name       string
	}{
		{"acb", "cba", "cdb", true, "Preset Case 1"},
		{"aaa", "a", "aab", false, "Preset Case 2"},
		{"aaa", "a", "aaaa", true, "Preset Case 3"},
		{"a", "a", "b", false, "Single letters - all 'a'"},
		{"a", "a", "a", true, "Single letters - all 'a'"},
		{"a", "b", "b", true, "Single letters - correct sum"},
		{"a", "b", "c", false, "Single letters - incorrect sum"},
		{"a", "a", "a", true, "All zero letters"},
		{"aaa", "a", "aaaa", true, "Zero sum"},
		{"aab", "c", "d", true, "Leading zeros in first word"},
		{"d", "aac", "f", true, "Leading zeros in second word"},
		{"aad", "aae", "aai", false, "Leading zeros in both words"},
		{strings.Repeat("j", 8), "a", strings.Repeat("j", 8), true, "long words"},
		{"jjjjjjji", "b", "jjjjjjjj", true, "biggest target word"},
		{"b", "jjjjjjji", "jjjjjjjj", true, "small first - big second"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.firstWord, testCase.secondWord, testCase.targetWord)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
