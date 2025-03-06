package problem2460

// Time complexity: O(2n). Space complexity: O(n).
func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i], nums[i+1] = nums[i]*2, 0
		}
	}
	ans := make([]int, len(nums))
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ans[idx] = nums[i]
			idx++
		}
	}
	return ans
}
