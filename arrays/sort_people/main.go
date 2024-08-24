package sortpeople

const factor = 1.25

func SortPeople(names []string, heights []int) []string {
	lengthOfData := len(heights)
	gap := lengthOfData
	swaps := true
	for gap > 1 || swaps {
		gap = max(1, int(float64(gap)/factor))
		swaps = false
		for i := 0; i < lengthOfData-gap; i++ {
			j := i + gap
			if heights[i] < heights[j] {
				heights[i], heights[j] = heights[j], heights[i]
				names[i], names[j] = names[j], names[i]
				swaps = true
			}
		}
	}
	return names
}
