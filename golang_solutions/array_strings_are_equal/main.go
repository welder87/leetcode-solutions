package arraystringsareequal

import "strings"


func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return true
}

func ArrayStringsAreEqual1(word1 []string, word2 []string) bool {
    return strings.Join(word1, "") == strings.Join(word2, "")
}
