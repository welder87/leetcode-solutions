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
	return right + 1 // or left
}

func searchInsertV1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)>>1
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
