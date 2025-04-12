package problem2089

import "sort"

// Time complexity: O(n log n) + O(log n) + O(log n) + O(m).
// Space complexity: O(n) + O(1) + O(1) + O(m).
// itmo pattern
func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	ans := make([]int, 0, len(nums))
	leftBoundary, rightBoundary := -1, -1
	left, right := -1, len(nums)
	for right > left+1 {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if right >= 0 && right < len(nums) && nums[right] == target {
		leftBoundary = right
		left = right
	} else {
		left = -1
	}
	right = len(nums)
	for right > left+1 {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}
	if left >= 0 && left < len(nums) && nums[left] == target && left != right {
		rightBoundary = left
	}
	if leftBoundary >= 0 && rightBoundary >= 0 {
		for i := leftBoundary; i <= rightBoundary; i++ {
			ans = append(ans, i)
		}
	}
	return ans
}

// Time complexity: O(2n) + O(m) + O(k).
// Space complexity: O(1) + O(m) + O(k).
func targetIndicesV1(nums []int, target int) []int {
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	counter := make([]int, maxVal+1)
	isFound := false
	for i := 0; i < len(nums); i++ {
		if !isFound && nums[i] == target {
			isFound = true
		}
		counter[nums[i]]++
	}
	if !isFound {
		return []int{}
	}
	cnt := 0
	for i := 1; i < len(counter); i++ {
		if i == target {
			break
		}
		cnt += counter[i]
	}
	rightBoundary := cnt + counter[target]
	ans := make([]int, 0, len(nums))
	for i := cnt; i < rightBoundary; i++ {
		ans = append(ans, i)
	}
	return ans
}
