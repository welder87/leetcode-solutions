package maximumwealth

func MaximumWealth(accounts [][]int) int {
	maxVal := 0
	current := 0
	for _, subArray := range accounts {
		current = 0
		for _, val := range subArray {
			current += val
		}
		if current > maxVal {
			maxVal = current
		}
	}
	return maxVal
}
