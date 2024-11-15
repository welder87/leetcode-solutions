package problem118

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		row := make([]int, 0, i+1)
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, result[i-1][j-1]+result[i-1][j])
		}
		row = append(row, 1)
		result = append(result, row)
	}
	return result
}
