package problem209

// Time complexity: O(n). Space complexity: O(1).
// modified itmo
func minSubArrayLen(target int, nums []int) int {
	left := 0
	sumVals := 0
	ans := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		sumVals += nums[right]
		for sumVals >= target {
			ans = min(ans, right-left+1)
			sumVals -= nums[left]
			left++
		}
	}
	if ans > len(nums) {
		return 0
	}
	return ans
}
