package problem2390

import (
	"testing"
)

func TestRemoveStars(t *testing.T) {
	testRemoveStars(t, removeStars)
}

func TestRemoveStarsV1(t *testing.T) {
	testRemoveStars(t, removeStarsV1)
}

func testRemoveStars(t *testing.T, fn func(string) string) {
	testCases := []struct {
		name string
		s    string
		ans  string
	}{
		{"Preset Case 1", "leet**cod*e", "lecoe"},
		{"Preset Case 2", "erase*****", ""},
		{"Single char then star", "a*", ""},
		{"Single char", "a", "a"},
		{"Without stars", "erase", "erase"},
		{"Single star at end", "abc*", "ab"},
		{"Single star in middle", "ab*c", "ac"},
		{"Two consecutive stars", "ab**c", "c"},
		{"Characters after multiple stars", "abc**de", "ade"},
		{"Three consecutive stars", "abc***", ""},
		{"Star after every character", "e*r*a*s*e*", ""},
		{"Stars at regular intervals", "a*a*a*a*a", "a"},
		{"Long string with stars", "hello*world***this*is*a**test****", "hellwothi"},
		{"Complex nested stars 1", "ab*cd**ef***gh", "gh"},
		{"Complex nested stars 2", "is**this****", ""},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s)
			if testCase.ans != ans {
				t.Errorf("got '%s', want '%s'", ans, testCase.ans)
			}
		})
	}
}
