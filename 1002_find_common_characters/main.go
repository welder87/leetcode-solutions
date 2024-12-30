package problem1002

func commonChars(words []string) []string {
	wordCount := len(words)
	symbolCount := 0
	counter := make(map[byte][]int)
	for i := 0; i < wordCount; i++ {
		word := words[i]
		for j := 0; j < len(word); j++ {
			symbol := word[j]
			if _, ok := counter[symbol]; !ok {
				counter[symbol] = make([]int, wordCount)
			}
			counter[symbol][i]++
			symbolCount++
		}
	}
	result := make([]string, 0, symbolCount)
	for symbol, counts := range counter {
		minCnt := counts[0]
		for i := 1; i < len(counts); i++ {
			if minCnt > counts[i] {
				minCnt = counts[i]
			}
		}
		if minCnt != 0 {
			for i := 0; i < minCnt; i++ {
				result = append(result, string(symbol))
			}
		}
	}
	return result
}
