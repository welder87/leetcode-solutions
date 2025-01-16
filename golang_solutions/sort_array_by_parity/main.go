package sortarraybyparity

func SortArrayByParity(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 == 0 {
			i++
		} else if nums[j]%2 != 0 {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	return nums
}
