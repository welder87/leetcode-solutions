package problem1880

// Time complexity: O(n + m + k). Space complexity: O(1).
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return calcNum(firstWord)+calcNum(secondWord) == calcNum(targetWord)
}

func calcNum(word string) rune {
	var m rune = 1
	var num rune
	for i := len(word) - 1; i >= 0; i-- {
		num += (rune(word[i]) - 'a') * m
		m *= 10
	}
	return num
}
