package problem551

import "strings"

// Time complexity: O(n). Space complexity: O(1).
func checkRecord(s string) bool {
	counterA, counterL, i := 0, 0, 0
	for i < len(s) {
		if s[i] != 'L' {
			if s[i] == 'A' {
				counterA++
			}
			i++
			continue
		}
		j := i
		for j < len(s) && s[j] == 'L' {
			j++
		}
		counterL = max(j-i, counterL)
		i = j
	}
	return counterA < 2 && counterL < 3
}

// Time complexity: O(n + n). Space complexity: O(1).
// Method: builtins.
// Solution: https://leetcode.doocs.org/en/lc/551 .
func checkRecordV1(s string) bool {
	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
}
