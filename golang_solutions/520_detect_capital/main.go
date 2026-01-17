package problem520

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
