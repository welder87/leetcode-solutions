package reverseprefix

func ReversePrefix(word string, ch byte) string {
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
