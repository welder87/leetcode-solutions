package problem3136

// Time complexity: O(n). Space complexity: O(1).
func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	vowel := false
	consonant := false
	for _, char := range word {
		isDigits := char >= '0' && char <= '9'
		isUpper := char >= 'A' && char <= 'Z'
		isLower := char >= 'a' && char <= 'z'
		if !isDigits && !isUpper && !isLower {
			return false
		}
		if isDigits {
			continue
		}
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowel = true
		} else {
			consonant = true
		}
	}
	return vowel && consonant
}
