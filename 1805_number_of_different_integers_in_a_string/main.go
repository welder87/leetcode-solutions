package problem1805

import "unicode"

func numDifferentIntegers(word string) int {
	if len(word) == 0 {
		return 0
	}
	nums := make(map[string]struct{}, len(word))
	num := make([]rune, 0, len(word))
	for idx, symbol := range word {
		if unicode.IsDigit(symbol) {
			if lastIndex := len(num) - 1; lastIndex >= 0 && len(num) == 1 && num[lastIndex] == rune('0') {
				num[lastIndex] = symbol
			} else {
				num = append(num, symbol)
			}
		} else if len(num) > 0 {
			nums[string(num)] = struct{}{}
			num = make([]rune, 0, len(word)-idx)
		}
	}
	if len(num) > 0 {
		nums[string(num)] = struct{}{}
	}
	return len(nums)
}
