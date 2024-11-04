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
