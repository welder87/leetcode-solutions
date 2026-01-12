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

// Time complexity: O(n+m). Space complexity: O(n+m).
// Solution: https://leetcode.doocs.org/en/lc/3174.
func clearDigitsV1(s string) string {
	stk := []byte{}
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, s[i])
		}
	}
	return string(stk)
}
