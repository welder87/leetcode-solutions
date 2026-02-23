package problem2109

// Time complexity O(2n + 2m), Space complexity O(2n + 2m).
func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	i, j := 0, 0
	for i < len(s) {
		if j < len(spaces) && i == spaces[j] {
			ans = append(ans, byte(' '))
			j++
		}
		ans = append(ans, s[i])
		i++
	}
	return string(ans)
}
