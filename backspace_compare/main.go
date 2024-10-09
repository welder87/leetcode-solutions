package backspacecompare

func BackspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	sharp := byte('#')
	for i >= 0 || j >= 0 {
		r1 := false
		if i >= 0 {
			r1 = s[i] == sharp
		}
		r2 := false
		if j >= 0 {
			r2 = t[j] == sharp
		}

		if r1 {
			counter := 1
			i--
			for i >= 0 && s[i] == sharp {
				counter++
				i--
			}
			i -= counter
		}
		if r2 {
			counter := 1
			j--
			for j >= 0 && t[j] == sharp {
				counter++
				j--
			}
			j -= counter
		}
		if r1 || r2 {
			continue
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	if i < 0 && j < 0 {
		return true
	}
	return false
}
