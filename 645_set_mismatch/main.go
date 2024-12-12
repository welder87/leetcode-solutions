package problem645

func findErrorNums(nums []int) []int {
	temp := make(map[int]int)
	for _, num := range nums {
		temp[num]++
	}
	result := []int{-1, -1}
	for i := 1; i <= len(nums); i++ {
		if cnt, ok := temp[i]; ok {
			if cnt > 1 {
				result[0] = i
			}
		} else {
			result[1] = i
		}
	}
	return result
}
