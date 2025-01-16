package getconcatenation


func GetConcatenation(nums []int) []int {
	srcLength := len(nums)
	newNums := make([]int, 2 *srcLength)
	var current int
	for i := 0; i < srcLength; i++ {
		current = nums[i]
		newNums[i] = current
		newNums[i + srcLength] = current
	}
	return newNums
}

func GetConcatenation1(nums []int) []int {
	srcLength := len(nums)
	for i := 0; i < srcLength; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}
