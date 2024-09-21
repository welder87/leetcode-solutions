package truncatesentence

import "strings"

func TruncateSentence(s string, k int) string {
	arr := strings.Split(s, " ")
	new_arr := []string{}
	for idx, sym := range arr {
		if idx >= k {
			break
		}
		new_arr = append(new_arr, string(sym))
	}
	return strings.Join(new_arr, " ")
}
