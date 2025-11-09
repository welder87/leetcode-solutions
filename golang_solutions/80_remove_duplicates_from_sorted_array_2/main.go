package problem80

// Time complexity: O(n). Space complexity: O(1).
// Solution: https://algo.monster/liteproblems/80
func removeDuplicates(nums []int) int {
	// Pointer to track the position for the next valid element
	writeIndex := 0
	// Iterate through each element in the array
	for i := range nums {
		// Allow element if:
		// 1. We have less than 2 elements (write_index < 2), OR
		// 2. Current element is different from the element 2 positions back
		if writeIndex < 2 || nums[i] != nums[writeIndex-2] {
			// Place the current element at the write position
			nums[writeIndex] = nums[i]
			// Move write pointer forward
			writeIndex += 1
		}
	}
	// Return the length of the modified array
	return writeIndex
}
