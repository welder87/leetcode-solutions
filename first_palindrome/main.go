package firstpalindrome

func FirstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		j := 0
		k := len(word) - 1
		isPalindrome := true
		for j <= k {
			if word[j] != word[k] {
				isPalindrome = false
				break
			}
			j++
			k--
		}
		if isPalindrome {
			return word
		}
	}
	return ""
}
