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

// Time complexity: O(2n). Space complexity: O(1).
func applyOperationsV1(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i], nums[i+1] = nums[i]*2, 0
		}
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i += 1
			}
		} else {
			i += 1
		}
		j += 1
	}
	return nums
}
