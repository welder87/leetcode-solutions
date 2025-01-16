package problem33

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	minimum := -1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] <= nums[len(nums)-1] {
			minimum = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if target > nums[len(nums)-1] {
		left = 0
		right = minimum
	} else {
		left = minimum
		right = len(nums) - 1
	}
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
