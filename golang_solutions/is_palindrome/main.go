package ispalindrome

import "unicode"

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !(('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') || ('0' <= s[i] && s[i] <= '9')) {
			i++
			continue
		}
		if !(('a' <= s[j] && s[j] <= 'z') || ('A' <= s[j] && s[j] <= 'Z') || ('0' <= s[j] && s[j] <= '9')) {
			j--
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func IsPalindrome1(s string) bool {
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
