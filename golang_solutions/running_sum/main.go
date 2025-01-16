package runningsum

func RunningSum(nums []int) []int {
	var prev int
	var current int
	for idx, val := range nums {
		current = val + prev
		nums[idx] = current
		prev = current
	}
	return nums
}
