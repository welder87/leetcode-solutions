package problem2942

// Time complexity: O(n * n). Space complexity: O(n)
func findWordsContaining(words []string, x byte) []int {
	lengthOfWords := len(words)
	slice := make([]int, 0, lengthOfWords)
	var lengthOfWord int
	var word string
	for i := 0; i < lengthOfWords; i++ {
		word = words[i]
		lengthOfWord = len(word)
		for j := 0; j < lengthOfWord; j++ {
			if word[j] == x {
				slice = append(slice, i)
				break
			}
		}
	}
	return slice
}
