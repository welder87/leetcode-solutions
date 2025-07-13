package problem3364

// Time complexity: O(n). Space complexity: O(1).
func minimumSumSubarray(nums []int, l int, r int) int {
	left := 0
	currentSum := 0
	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		diff := right - left + 1
		for diff <= r && diff >= l {
			currentSum -= nums[left]
		}
	}
	return 0
}
