package problem551

// Time complexity: O(n). Space complexity: O(1).
func checkRecord(s string) bool {
	counterA, counterL, i := 0, 0, 0
	for i < len(s) {
		if s[i] != 'L' {
			if s[i] == 'A' {
				counterA++
			}
			i++
			continue
		}
		j := i
		for j < len(s) && s[j] == 'L' {
			j++
		}
		counterL = max(j-i, counterL)
		i = j
	}
	return counterA < 2 && counterL < 3
}
