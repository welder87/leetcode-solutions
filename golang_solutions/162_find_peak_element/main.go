package problem162

// Time complexity: O(log n). Space complexity: O(1).
// modified leetcode pattern
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
