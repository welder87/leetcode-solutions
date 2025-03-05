package problem26

// Time complexity: O(n). Space complexity: O(1).
func removeDuplicates(nums []int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if j == 0 || nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// Time complexity: O(n). Space complexity: O(1).
func removeDuplicatesV1(nums []int) int {
	i, j := 0, 1
	for i < len(nums) {
		for j < len(nums) && nums[i] == nums[j] {
			j += 1
		}
		if j < len(nums) {
			nums[i+1] = nums[j]
		} else {
			break
		}
		i++
	}
	return i + 1
}
