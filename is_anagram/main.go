package isanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter1 := map[byte]int{}
	counter2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if val, isFound := counter1[s[i]]; isFound {
			counter1[s[i]] = val + 1
		} else {
			counter1[s[i]] = 1
		}
		if val, isFound := counter2[t[i]]; isFound {
			counter2[t[i]] = val + 1
		} else {
			counter2[t[i]] = 1
		}
	}
	for key, val1 := range counter1 {
		if val2, isFound := counter2[key]; !isFound || val1 != val2 {
			return false
		}
	}
	return true
}

// func caclCount(elem byte, counter *map[byte]int) {
// 	if val, isFound := counter[elem]; isFound {
// 		counter[elem] = val + 1
// 	} else {
// 		counter[elem] = 1
// 	}
// }
