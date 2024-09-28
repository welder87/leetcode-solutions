package reverseonlyletters

import "unicode"

func ReverseOnlyLetters(s string) string {
	i := 0
	j := len(s) - 1
	newS := make([]byte, len(s))
	for i < j {
		leftOk := unicode.IsLetter(rune(s[i]))
		rightOk := unicode.IsLetter(rune(s[j]))
		if leftOk && rightOk {
			newS[i], newS[j] = s[j], s[i]
			i++
			j--
		}
		if !leftOk {
			newS[i] = s[i]
			i++
		}
		if !rightOk {
			newS[j] = s[j]
			j--
		}
	}
	if i == j {
		newS[i] = s[i]
	}
	return string(newS)
}
