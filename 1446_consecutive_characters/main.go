package problem1446

func maxPower(s string) int {
	prev := s[0]
	counter := 1
	maxCounter := 0
	for i := 1; i < len(s); i++ {
		if prev == s[i] {
			counter++
		} else {
			prev = s[i]
			if counter > maxCounter {
				maxCounter = counter
			}
			counter = 1
		}
	}
	if counter > maxCounter {
		maxCounter = counter
	}
	return maxCounter
}
