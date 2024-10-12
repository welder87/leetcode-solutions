package lengthoflastword

func LengthOfLastWord(s string) int {
	whitespace := byte(' ')
	counter := 0
	isFound := false
	for i := len(s) - 1; i >= 0; i-- {
		if isFound {
			if s[i] == whitespace {
				return counter
			} else {
				counter++
			}
		} else {
			if s[i] != whitespace {
				isFound = true
				counter++
			}
		}
	}
	return counter
}
