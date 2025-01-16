package problem1876

func countGoodSubstrings(s string) int {
	lenght := 3
	if len(s) < lenght {
		return 0
	}
	counter := 0
	symbols := make(map[byte]int, lenght)
	for i := 0; i < lenght; i++ {
		symbols[s[i]]++
	}
	if len(symbols) == lenght {
		counter++
	}
	index := 0
	if cnt, ok := symbols[s[index]]; ok && cnt > 1 {
		symbols[s[index]]--
	} else {
		delete(symbols, s[index])
	}
	for i := lenght; i < len(s); i++ {
		symbols[s[i]]++
		if len(symbols) == lenght {
			counter++
		}
		index := i - lenght + 1
		if cnt, ok := symbols[s[index]]; ok && cnt > 1 {
			symbols[s[index]]--
		} else {
			delete(symbols, s[index])
		}
	}
	return counter
}
