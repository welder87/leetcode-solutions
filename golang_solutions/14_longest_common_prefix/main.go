package problem14

import "strings"

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

// Time complexity: O(sum(char)). Space complexity: O(1).
// Method: Horizontal scanning
// Solution: https://leetcode.com/problems/longest-common-prefix/editorial/#approach-1-horizontal-scanning
func longestCommonPrefixV1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// Time complexity: O(sum(char)). Space complexity: O(1).
// Method: Vertical scanning
// Solution: https://leetcode.com/problems/longest-common-prefix/editorial/#approach-2-vertical-scanning
func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// Time complexity: O(sum(char)). Space complexity: O(1).
func longestCommonPrefixV3(strs []string) string {
	ans := []byte(strs[0])
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
	return string(ans)
}
