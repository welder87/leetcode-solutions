package problem136

func singleNumber(nums []int) int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}
	for val, cnt := range counter {
		if cnt == 1 {
			return val
		}
	}
	return 0
}
