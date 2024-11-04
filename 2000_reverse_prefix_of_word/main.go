package problem2000

func ReversePrefix(word string, ch byte) string {
	if len(word) == 0 || len(word) == 1 {
		return word
	}
	i := 0
	j := -1
	for i < len(word) {
		if word[i] == ch {
			j = i
			break
		}
		i++
	}
	if j == -1 || j == len(word) {
		return word
	}
	i = 0
	newWord := []byte(word)
	for i < j {
		newWord[i], newWord[j] = newWord[j], newWord[i]
		i++
		j--
	}
	return string(newWord)
}

func ReversePrefixV1(word string, ch byte) string {
	idx := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			idx = i
			break
		}
	}
	if idx == -1 || idx == len(word) {
		return word
	}
	newString := make([]byte, len(word))
	i := idx
	j := idx + 1
	for {
		leftOk := i >= 0
		rightOk := j < len(word)
		if !leftOk && !rightOk {
			return string(newString)
		}
		if leftOk {
			newString[idx-i] = word[i]
			i--
		}
		if rightOk {
			newString[j] = word[j]
			j++
		}
	}
}
