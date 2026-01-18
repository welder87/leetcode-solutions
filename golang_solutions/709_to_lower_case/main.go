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
