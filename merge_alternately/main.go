package mergealternately

func MergeAlternately(word1 string, word2 string) string {
	newStringArray := make([]byte, 0, len(word1)+len(word2))
	i := 0
	j := 0
	for {
		ok1 := i < len(word1)
		ok2 := j < len(word2)
		if !ok1 && !ok2 {
			return string(newStringArray)
		}
		if ok1 {
			newStringArray = append(newStringArray, word1[i])
			i++
		}
		if ok2 {
			newStringArray = append(newStringArray, word2[j])
			j++
		}
	}
}
