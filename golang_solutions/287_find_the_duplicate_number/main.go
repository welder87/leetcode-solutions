package problem287

import "sort"

// Time complexity: O(n + m). Space complexity: O(m).
func findDuplicate(nums []int) int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}
	for num, cnt := range counter {
		if cnt > 1 {
			return num
		}
	}
	return 0
}

// Time complexity: O(n^2). Space complexity: O(1).
func findDuplicateV1(nums []int) int {
	for _, num1 := range nums {
		cnt := 0
		for _, num2 := range nums {
			if num1 == num2 {
				cnt += 1
			}
		}
		if cnt > 1 {
			return num1
		}
	}
	return 0
}

// Time complexity: O(n log n + n). Space complexity: O(n log n).
func findDuplicateV2(nums []int) int {
	sort.Ints(nums)
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == nums[j] {
			return nums[i]
		}
		i++
		j++
	}
	return 0
}
