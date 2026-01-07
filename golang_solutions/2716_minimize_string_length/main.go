package problem2716

// Time complexity: O(n). Space complexity: O(m).
func minimizedStringLength(s string) int {
	uniqueSym := make(map[byte]struct{}, len(s))
	for i := range s {
		if _, ok := uniqueSym[s[i]]; !ok {
			uniqueSym[s[i]] = struct{}{}
		}
	}
	return len(uniqueSym)
}

// Time complexity: O(n). Space complexity: O(m).
func minimizedStringLengthV1(s string) int {
	uniqueSym := make(map[byte]struct{}, len(s))
	for i := range s {
		uniqueSym[s[i]] = struct{}{}
	}
	return len(uniqueSym)
}

// Time complexity: O(n). Space complexity: O(m).
// Solution: https://leetcode.doocs.org/en/lc/2716/.
func minimizedStringLengthV2(s string) int {
	ss := map[rune]struct{}{}
	for _, c := range s {
		ss[c] = struct{}{}
	}
	return len(ss)
}
