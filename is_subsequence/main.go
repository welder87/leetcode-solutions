package issubsequence

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	j := 0
	counter := 0
	for {
		if i >= len(s) || j >= len(t) {
			return false
		}
		if s[i] == t[j] {
			counter++
			if counter == len(s) {
				return true
			}
			i++
		}
		j++
	}
}
