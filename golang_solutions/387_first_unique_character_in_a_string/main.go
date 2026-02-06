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

// Time complexity: O(n). Space complexity: O(1).
// Solution: https://leetcode.doocs.org/en/lc/387
func firstUniqCharV1(s string) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, c := range s {
		if cnt[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
