package problem153

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	ans := -1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] <= nums[len(nums)-1] {
			ans = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return nums[ans]
}
