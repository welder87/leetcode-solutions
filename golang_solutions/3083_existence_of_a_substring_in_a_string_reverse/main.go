package problem3083

// Time complexity: O(n+n). Space complexity: O(m).
// Method: Pair As Byte Array.
func isSubstringPresent(s string) bool {
	uniquePair := make(map[[2]byte]struct{}, len(s))
	for i := range len(s) - 1 {
		uniquePair[[2]byte{s[i], s[i+1]}] = struct{}{}
	}
	for i := len(s) - 1; i > 0; i-- {
		if _, ok := uniquePair[[2]byte{s[i], s[i-1]}]; ok {
			return true
		}
	}
	return false
}
