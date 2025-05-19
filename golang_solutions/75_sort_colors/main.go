package problem75

// Time complexity: O(n*n). Space complexity: O(1).
// Selection sort
func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		if nums[minIdx] < nums[i] {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

// Time complexity: O(n*n). Space complexity: O(1).
// Bubble sort
func sortColorsV1(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
