package problem283

// Time complexity: O(n). Space complexity: O(1).
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := 0
	j := 1
	for i < (len(nums)-1) && j < len(nums) {
		if nums[i] != 0 && nums[j] != 0 {
			i++
			j++
		} else if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[i] == 0 && nums[j] == 0 {
			j++
		} else {
			i++
		}
	}
}

// Time complexity: O(n). Space complexity: O(1).
func moveZeroesV1(nums []int) {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		} else if nums[i] != 0 {
			i += 1
		}
		j += 1
	}
}
