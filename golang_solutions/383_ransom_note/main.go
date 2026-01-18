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

// Time complexity: O(n + m). Space complexity: O(k).
// Solution: https://leetcode.doocs.org/en/lc/383
func canConstructV1(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for _, c := range magazine {
		cnt[c-'a']++
	}
	for _, c := range ransomNote {
		cnt[c-'a']--
		if cnt[c-'a'] < 0 {
			return false
		}
	}
	return true
}
