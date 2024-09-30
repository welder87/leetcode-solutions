package strstr

func StrStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		}
		return -1
	}
	lastIndex := len(needle) - 1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := i
		n := 0
		for n < len(needle) {
			if haystack[j] != needle[n] {
				break
			}
			if lastIndex == n {
				return i
			}
			j++
			n++
		}
	}
	return -1
}
