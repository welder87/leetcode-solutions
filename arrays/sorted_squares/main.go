package sortedsquares

func SortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	for k := len(result) - 1; k >= 0; k-- {
		leftSquare := nums[i] * nums[i]
		rightSquare := nums[j] * nums[j]
		if leftSquare > rightSquare {
			result[k] = leftSquare
			i++
		} else {
			result[k] = rightSquare
			j--
		}
	}
	return result
}
