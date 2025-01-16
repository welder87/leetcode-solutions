package problem151

func reverseWords(s string) string {
	result := make([]byte, 0, len(s))
	start := -1
	end := -1
	i := len(s) - 1
	for i >= 0 {
		if s[i] != byte(' ') && start < 0 {
			start = i
		} else if s[i] == byte(' ') && start > 0 && end < 0 {
			end = i + 1
		}
		if start > 0 && end >= 0 {
			for j := end; j < start+1; j++ {
				result = append(result, s[j])
			}
			result = append(result, byte(' '))
			start = -1
			end = -1
		}
		i--
	}
	if start >= 0 && end < 0 {
		end = i + 1
	}
	if start >= 0 && end >= 0 {
		for j := end; j < start+1; j++ {
			result = append(result, s[j])
		}
		result = append(result, byte(' '))
	}
	return string(result[:len(result)-1])
}
