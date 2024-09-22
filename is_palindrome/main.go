package ispalindrome

import "unicode"

func IsPalindrome(s string) bool {
	vectorOfRune := make([]rune, 0, len(s))
	for _, val := range s {
		if ('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9') {
			vectorOfRune = append(vectorOfRune, unicode.ToLower(val))
		}
	}
	i := 0
	j := len(vectorOfRune) - 1
	for i < j {
		if vectorOfRune[i] != vectorOfRune[j] {
			return false
		}
		i++
		j--
	}
	return true
}
