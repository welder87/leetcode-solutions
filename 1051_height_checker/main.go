package problem1051

func heightChecker(heights []int) int {
	maxHeight := -1
	for _, height := range heights {
		if height > maxHeight {
			maxHeight = height
		}
	}
	counter := make([]int, maxHeight+1)
	for _, height := range heights {
		counter[height]++
	}
	for i := 1; i < maxHeight+1; i++ {
		counter[i] += counter[i-1]
	}
	sortedHeights := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		counter[heights[i]]--
		sortedHeights[counter[heights[i]]] = heights[i]
	}
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			count++
		}
	}
	return count
}
