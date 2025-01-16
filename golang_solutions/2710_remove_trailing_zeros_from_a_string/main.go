package problem2710

func removeTrailingZeros(num string) string {
	i := len(num) - 1
	for i >= 0 {
		if num[i] != byte('0') {
			break
		} else {
			i--
		}
	}
	return num[:i+1]
}
