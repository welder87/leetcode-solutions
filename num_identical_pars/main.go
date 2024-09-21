package numidenticalpars

func NumIdenticalPairs(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		val, isFound := counter[num]
		if isFound {
			counter[num] = val + 1
		} else {
			counter[num] = 1
		}
	}
	pairCount := 0
	for _, value := range counter {
		if value > 1 {
			pairCount += value * (value - 1) / 2
		}
	}
	return pairCount
}
