package problem3794

// Time complexity O(n + n), space complexity O(n + n)
// Method: Two-pointer approach (converging pointers).
func reversePrefix(s string, k int) string {
	ans := []rune(s)
	i, j := 0, k-1
	for i < j {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	return string(ans)
}
