package strstr

// Алгоритм Кнута-Морриса-Пратта
// O(n)
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
	return calcIndex(haystack, needle, makeArrayOfSuffix(needle))
}

func makeArrayOfSuffix(needle string) []int {
	pi := make([]int, len(needle))
	j := 0
	i := 1
	for i < len(needle) {
		if needle[j] == needle[i] {
			pi[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				pi[i] = 0
				i++
			} else {
				j = pi[j-1]
			}
		}
	}
	return pi
}

func calcIndex(haystack string, needle string, pi []int) int {
	i := 0
	j := 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		} else {
			if j > 0 {
				j = pi[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

// Базовый алгоритм O(n^2)
// O(n^2)
func StrStr1(haystack string, needle string) int {
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
