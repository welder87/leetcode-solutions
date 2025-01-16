package problem1

func twoSum(nums []int, target int) []int {
	calcs := make(map[int]int, len(nums))
	for idx, val := range nums {
		item := target - val
		if index, ok := calcs[item]; ok {
			return []int{index, idx}
		}
		calcs[val] = idx
	}
	return nil
}

func twoSumV1(nums []int, target int) []int {
	arr := []int{-1, -1}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				arr[0] = i
				arr[1] = j
				return arr
			}
		}
	}
	return nil
}
