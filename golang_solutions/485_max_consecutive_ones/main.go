package problem485

func findMaxConsecutiveOnes(nums []int) int {
	maxCounter := 0
	counter := 0
	for _, num := range nums {
		if num == 1 {
			counter++
		} else {
			if counter > maxCounter {
				maxCounter = counter
			}
			counter = 0
		}
	}
	if counter > maxCounter {
		maxCounter = counter
	}
	return maxCounter
}
