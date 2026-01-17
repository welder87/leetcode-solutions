package problem520

import "unicode"

// Time complexity: O(n). Space complexity: O(1).
func detectCapitalUse(word string) bool {
	first, uppperCaseCounter := false, 0
	for idx, char := range word {
		if idx == 0 && (char >= 'A' && char <= 'Z') {
			first = true
		}
		if char >= 'A' && char <= 'Z' {
			uppperCaseCounter++
		}
	}
	return uppperCaseCounter == len(word) || uppperCaseCounter == 0 || first && uppperCaseCounter == 1
}

// Time complexity: O(n). Space complexity: O(1).
// Solution: https://leetcode.doocs.org/en/lc/520/
func detectCapitalUseV1(word string) bool {
	cnt := 0
	for _, c := range word {
		if unicode.IsUpper(c) {
			cnt++
		}
	}
	return cnt == 0 || cnt == len(word) || (cnt == 1 && unicode.IsUpper(rune(word[0])))
}
