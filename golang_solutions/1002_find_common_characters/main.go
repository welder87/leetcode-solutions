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

func commonCharsV1(words []string) []string {
	wordCount := len(words)
	mainCounter := make(map[byte]int, len(words[0]))
	for j := 0; j < len(words[0]); j++ {
		mainCounter[words[0][j]]++
	}
	for i := 1; i < wordCount; i++ {
		counter := make(map[byte]int, len(words[i]))
		for j := 0; j < len(words[i]); j++ {
			counter[words[i][j]]++
		}
		for symbol, mainCnt := range mainCounter {
			if cnt, ok := counter[symbol]; ok {
				mainCounter[symbol] = min(mainCnt, cnt)
			} else {
				mainCounter[symbol] = 0
			}
		}
	}
	result := make([]string, 0, len(words[0]))
	for symbol, cnt := range mainCounter {
		for i := 0; i < cnt; i++ {
			result = append(result, string(symbol))
		}
	}
	return result
}
