package problem1920

func buildArray(nums []int) []int {
	lengthOfNums := len(nums)
	new_nums := make([]int, lengthOfNums)
	for i := 0; i < lengthOfNums; i++ {
		new_nums[i] = nums[nums[i]]
	}
	return new_nums
}
