package problem14

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testLongestCommonPrefix(t, longestCommonPrefix)
}

func TestLongestCommonPrefixV1(t *testing.T) {
	testLongestCommonPrefix(t, longestCommonPrefixV1)
}

func TestLongestCommonPrefixV2(t *testing.T) {
	testLongestCommonPrefix(t, longestCommonPrefixV2)
}

func TestLongestCommonPrefixV3(t *testing.T) {
	testLongestCommonPrefix(t, longestCommonPrefixV3)
}

func testLongestCommonPrefix(t *testing.T, fn func([]string) string) {
	testCases := []struct {
		strs []string
		ans  string
		name string
	}{
		{[]string{"flower", "flow", "flight"}, "fl", "Preset Case 1"},
		{[]string{"dog", "racecar", "car"}, "", "Preset Case 2"},
		{[]string{"flower"}, "flower", "Single non empty word"},
		{[]string{""}, "", "Single empty word"},
		{[]string{"", ""}, "", "Multiple empty strings"},
		{[]string{"a"}, "a", "Single character"},
		{[]string{"a", "a"}, "a", "Same single characters"},
		{[]string{"a", "b"}, "", "Different single characters"},
		{[]string{"abc", "abc", "abc"}, "abc", "All identical"},
		{[]string{"flower", "flow", "fright"}, "f", "Single character in end"},
		{[]string{"fright", "flower", "flow"}, "f", "Single character in start"},
		{[]string{"flower", "flowe", "flow"}, "flow", "Full word in end"},
		{[]string{"flow", "flower", "flowe"}, "flow", "Full word in start"},
		{[]string{"flow", "flower", ""}, "", "Empty word in end"},
		{[]string{"", "flower", "flowe"}, "", "Empty word in start"},
		{[]string{"throne", "throne"}, "throne", "Two same words"},
		{
			[]string{
				strings.Repeat("a", 1000),
				strings.Repeat("a", 1000),
				strings.Repeat("a", 1000),
			},
			strings.Repeat("a", 1000),
			"Performance Case 1",
		},
		{
			[]string{
				"pre" + strings.Repeat("x", 1000),
				"pre" + strings.Repeat("y", 1000),
			},
			"pre",
			"Performance Case 2",
		},
		{
			[]string{strings.Repeat("abcdefgh", 100), strings.Repeat("abcdefgh", 100)},
			strings.Repeat("abcdefgh", 100),
			"Performance Case 3",
		},
		{
			[]string{"interspecies", "interstellar", "interstate"},
			"inters",
			"Success Case 1",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.strs); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
