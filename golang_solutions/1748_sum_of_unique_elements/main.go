package problem1748

func sumOfUnique(nums []int) int {
	counter := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	ans := 0
	for num, cnt := range counter {
		if cnt == 1 {
			ans += num
		}
	}
	return ans
}
