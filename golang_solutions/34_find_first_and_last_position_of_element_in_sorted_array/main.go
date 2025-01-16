package problem34

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{getBoundary(nums, target, true), getBoundary(nums, target, false)}
}

func getBoundary(nums []int, target int, isLeft bool) int {
	left := 0
	right := len(nums) - 1
	ans := -1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] == target {
			ans = middle
			if isLeft {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return ans
}
