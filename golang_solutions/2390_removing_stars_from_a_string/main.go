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

// Time complexity: O(n + m). Space complexity: O(n + m).
// Solution: https://leetcode.com/problems/removing-stars-from-a-string/solutions/7244290/clean-simple-beats-9587-easy-explanation-m7e5/
func removeStarsV1(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c != '*' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
