package problem2129

import "strings"

// Time complexity O(n + n + n), Space complexity O(n + n)
func capitalizeTitle(title string) string {
	ans := []rune(title)
	start := 0
	for i := range ans {
		if ans[i] == ' ' {
			if i-start > 2 {
				ans[start] = ans[start] - 32
			}
			start = i + 1
		} else if 'A' <= ans[i] && ans[i] <= 'Z' {
			ans[i] = ans[i] + 32
		}
	}
	if len(ans)-start > 2 {
		ans[start] = ans[start] - 32
	}
	return string(ans)
}

// Time complexity O(n + n + n), Space complexity O(n + n)
// Solution: https://leetcode.doocs.org/en/lc/2129 .
func capitalizeTitleV1(title string) string {
	title = strings.ToLower(title)
	words := strings.Split(title, " ")
	for i, s := range words {
		if len(s) > 2 {
			words[i] = strings.Title(s)
		}
	}
	return strings.Join(words, " ")
}
