package problem2733

import "sort"

// Time complexity: O(n log n). Space complexity: O(1).
func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	return nums[1]
}

// Time complexity: O(3n). Space complexity: O(m).
func findNonMinOrMaxV1(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	sortedNums := make([]bool, (maxNum + 1))
	for i := 0; i < len(nums); i++ {
		sortedNums[nums[i]] = true
	}
	counter := 0
	for i := 0; i < len(sortedNums); i++ {
		if sortedNums[i] {
			counter += 1
		}
		if counter == 2 {
			return i
		}
	}
	return -1
}

// Time complexity: O(2n). Space complexity: O(1).
func findNonMinOrMaxV2(nums []int) int {
	minNum, maxNum := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
		if minNum > nums[i] {
			minNum = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if minNum < nums[i] && nums[i] < maxNum {
			return nums[i]
		}
	}
	return -1
}
