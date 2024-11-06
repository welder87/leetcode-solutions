package problem1961

func isPrefixString(s string, words []string) bool {
	total := 0
	result := make([]byte, 0, len(s))
	for _, word := range words {
		total += len(word)
		result = append(result, []byte(word)...)
		if total == len(s) && string(result) == s {
			return true
		}
	}
	return false
}
