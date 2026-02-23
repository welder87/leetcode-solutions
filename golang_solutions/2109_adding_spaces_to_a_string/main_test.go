package problem2109

import (
	"testing"
)

func TestAddSpaces(t *testing.T) {
	testAddSpaces(t, addSpaces)
}

func testAddSpaces(t *testing.T, fn func(string, []int) string) {
	testCases := []struct {
		s      string
		spaces []int
		ans    string
		name   string
	}{
		{
			"LeetcodeHelpsMeLearn",
			[]int{8, 13, 15},
			"Leetcode Helps Me Learn",
			"Preset Case 1",
		},
		{
			"icodeinpython",
			[]int{1, 5, 7, 9},
			"i code in py thon",
			"Preset Case 2",
		},
		{
			"spacing",
			[]int{0, 1, 2, 3, 4, 5, 6},
			" s p a c i n g",
			"Preset Case 3",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if ans := fn(testCase.s, testCase.spaces); testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
