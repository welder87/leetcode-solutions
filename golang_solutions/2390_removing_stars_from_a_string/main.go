package problem2390

// Time complexity: O(n + m). Space complexity: O(n + m).
func removeStars(s string) string {
	ans := []rune{}
	for _, sym := range s {
		if sym != '*' {
			ans = append(ans, sym)
		} else if sym == '*' && len(ans) > 0 {
			ans = ans[:len(ans)-1]
		}
	}
	return string(ans)
}
