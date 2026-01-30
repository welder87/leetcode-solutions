package problem14

// Time complexity: O(sum(char)). Space complexity: O(1).
func longestCommonPrefix(strs []string) string {
	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(ans) || j < len(strs[i]) {
			if j == len(strs[i]) || j < len(ans) && ans[j] != strs[i][j] {
				ans = ans[:j]
				break
			}
			j++
		}
	}
	return ans
}
