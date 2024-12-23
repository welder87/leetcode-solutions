package problem35

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return right + 1
}
