package mostwordsfound

import (
	"strings"
)

func MostWordsFound(sentences []string) int {
	var maxCount int
	one_sentence := strings.Join(sentences, "0")
	counter := 0
	for _, val := range one_sentence {
		if val == 48 {
			if counter > maxCount {
				maxCount = counter
			}
			counter = 0
		} else if val == 32 {
			counter++
		}
	}
	if counter > maxCount {
		maxCount = counter
	}
	return maxCount + 1
}

func MostWordsFound1(sentences []string) int {
	var maxCount int
	for _, sentence := range sentences {
		counter := 0
		for _, val := range sentence {
			if val == 32 {
				counter++
			}
		}
		if counter > maxCount {
			maxCount = counter
		}
	}
	return maxCount + 1
}
