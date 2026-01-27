package problem3110

// Time complexity: O(n). Space complexity: O(1).
func scoreOfString(s string) int {
	sum_ := 0
	for i := 0; i < len(s)-1; i++ {
		sum_ += abs(rune(s[i]) - rune(s[i+1]))
	}
	return sum_
}

func abs(x rune) int {
	if x < 0 {
		return int(-x)
	}
	return int(x)
}
