package problem1929

// Time complexity: O(n). Space complexity: O(2n).
func getConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+len(nums)] = nums[i]
	}
	return ans
}

// Time complexity: O(n). Space complexity: O(2n).
func getConcatenationV1(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
