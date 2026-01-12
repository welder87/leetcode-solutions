package problem3174

import "unicode"

// Time complexity: O(n+m). Space complexity: O(n+m).
func clearDigits(s string) string {
	ans := []rune{}
	for _, sym := range s {
		if unicode.IsDigit(sym) {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, sym)
		}
	}
	return string(ans)
}
