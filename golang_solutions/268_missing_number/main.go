package problem268

import "sort"

// Time complexity: O(n^2). Space complexity: O(m).
func missingNumber(nums []int) int {
	maxIdx := len(nums) + 1
	for i := 0; i < maxIdx; i++ {
		isFound := false
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				isFound = true
				break
			}
		}
		if !isFound {
			return i
		}
	}
	return -1
}

// Time complexity: O(2n). Space complexity: O(n).
func missingNumberV1(nums []int) int {
	uniqueNums := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		uniqueNums[nums[i]] = struct{}{}
	}
	maxIdx := len(nums) + 1
	for i := 0; i < maxIdx; i++ {
		if _, ok := uniqueNums[i]; !ok {
			return i
		}
	}
	return -1
}

// Time complexity: O(n log n + n + 1). Space complexity: O(n log n).
func missingNumberV2(nums []int) int {
	sort.Ints(nums)
	maxIdx := len(nums) + 1
	for i := 0; i < maxIdx; i++ {
		if len(nums) == i || i != nums[i] {
			return i
		}
	}
	return -1
}

// Time complexity: O(2n). Space complexity: O(1).
func missingNumberV3(nums []int) int {
	current := 0
	for i := 0; i < len(nums); i++ {
		current += nums[i]
	}
	reference := 0
	maxIdx := len(nums) + 1
	for i := 0; i < maxIdx; i++ {
		reference += i
	}
	if val := reference - current; val < 0 {
		return val * -1
	} else {
		return val
	}
}

// Time complexity: O(2n). Space complexity: O(1).
func missingNumberV4(nums []int) int {
	current := 0
	for i := 0; i < len(nums); i++ {
		current ^= nums[i]
	}
	reference := 0
	maxIdx := len(nums) + 1
	for i := 0; i < maxIdx; i++ {
		reference ^= i
	}
	return current ^ reference
}
