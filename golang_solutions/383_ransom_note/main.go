package problem383

// Time complexity: O(n + m + n). Space complexity: O(n + m).
func canConstruct(ransomNote string, magazine string) bool {
	counter1 := make(map[rune]int, len(ransomNote))
	counter2 := make(map[rune]int, len(magazine))
	for _, sym := range ransomNote {
		counter1[sym]++
	}
	for _, sym := range magazine {
		counter2[sym]++
	}
	for sym, cnt1 := range counter1 {
		if cnt2, ok := counter2[sym]; !ok || cnt2 < cnt1 {
			return false
		}
	}
	return true
}
