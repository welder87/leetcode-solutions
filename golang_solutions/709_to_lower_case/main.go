package problem704

import (
	"strings"
	"unicode"
)

// Time complexity: O(n + m). Space complexity: O(n + m).
func toLowerCase(s string) string {
	ans := make([]rune, 0, len(s))
	for _, char := range s {
		ans = append(ans, unicode.ToLower(char))
	}
	return string(ans)
}

// Time complexity: O(n + m). Space complexity: O(n + m).
func toLowerCaseV1(s string) string {
	return strings.ToLower(s)
}

// Time complexity: O(n + m). Space complexity: O(n + m).
// Solution:
// https://leetcode.com/problems/to-lower-case/solutions/7249092/simple-golang-solution-beats-100-by-srgl-jy2b/
func toLowerCaseV2(s string) string {
	res := make([]rune, len(s))
	for i, r := range s {
		res[i] = r
		if r <= 'Z' && r >= 'A' {
			res[i] = rune(r + 32)
		}
	}
	return string(res)
}

// Time complexity: O(n + m). Space complexity: O(n + m).
// Solution: https://leetcode.doocs.org/en/lc/709
func toLowerCaseV3(s string) string {
	cs := []byte(s)
	for i, c := range cs {
		if c >= 'A' && c <= 'Z' {
			cs[i] |= 32
		}
	}
	return string(cs)
}
