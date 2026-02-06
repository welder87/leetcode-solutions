package problem389

// Time complexity: O(n). Space complexity: O(1).
func firstUniqChar(s string) int {
	counter := make(map[rune]int, len(s))
	for _, sym := range s {
		counter[sym]++
	}
	for idx, sym := range s {
		if counter[sym] == 1 {
			return idx
		}
	}
	return -1
}
