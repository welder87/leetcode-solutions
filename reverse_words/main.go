package reversewords

import "strings"

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	var space byte = 32
	reversedWords := make([]byte, 0, len(s))
	lastIndex := len(words) - 1
	for index, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			reversedWords = append(reversedWords, word[i])
		}
		if index < lastIndex {
			reversedWords = append(reversedWords, space)
		}
	}
	return string(reversedWords)
}

// func ReverseString(s []byte) {
// 	i := 0
// 	j := len(s) - 1
// 	for i < j {
// 		s[i], s[j] = s[j], s[i]
// 		i++
// 		j--
// 	}
// }
