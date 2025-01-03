package problem2248

import "sort"

func intersection(nums [][]int) []int {
	idxWithMinLength := -1
	for idx := range nums {
		if idxWithMinLength == -1 || len(nums[idxWithMinLength]) > len(nums[idx]) {
			idxWithMinLength = idx
		}
	}
	existing := make(map[int]struct{}, len(nums[idxWithMinLength]))
	for _, num := range nums[idxWithMinLength] {
		existing[num] = struct{}{}
	}
	for idx, arr := range nums {
		if idx == idxWithMinLength {
			continue
		}
		current := make(map[int]struct{}, len(arr))
		for _, num := range arr {
			if _, ok := existing[num]; ok {
				current[num] = struct{}{}
			}
		}
		for num := range existing {
			if _, ok := current[num]; !ok {
				delete(existing, num)
			}
		}
	}
	result := make([]int, 0, len(nums[0]))
	for num := range existing {
		result = append(result, num)
	}
	sort.Ints(result)
	return result
}
