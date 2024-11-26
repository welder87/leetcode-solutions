package problem119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	prevRow := &[]int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		currentRow := make([]int, 0, i+1)
		currentRow = append(currentRow, 1)
		for j, k := 0, 1; k < len(*prevRow); j, k = j+1, k+1 {
			currentRow = append(currentRow, (*prevRow)[j]+(*prevRow)[k])
		}
		currentRow = append(currentRow, 1)
		prevRow = &currentRow
	}
	return *prevRow
}
