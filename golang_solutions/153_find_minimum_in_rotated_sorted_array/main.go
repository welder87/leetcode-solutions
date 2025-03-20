package problem153

// Time complexity: O(log n). Space complexity: O(1).
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
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

// Time complexity: O(log n). Space complexity: O(1).
func findMinV1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
