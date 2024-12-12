package problem645

import "sort"

func findErrorNums(nums []int) []int {
	temp := make(map[int]int)
	for _, num := range nums {
		temp[num]++
	}
	result := []int{-1, -1}
	for i := 1; i <= len(nums); i++ {
		if cnt, ok := temp[i]; ok {
			if cnt > 1 {
				result[0] = i
			}
		} else {
			result[1] = i
		}
	}
	return result
}

func findErrorNumsV1(nums []int) []int {
	sort.Ints(nums)
	result := []int{-1, -1}
	sum1 := 0
	sum2 := nums[0]
	i := 1
	for i < len(nums) {
		if nums[i]-nums[i-1] == 0 {
			result[0] = nums[i]
		}
		sum1 += i
		sum2 += nums[i]
		i++
	}
	sum1 += i
	result[1] = (sum1 - sum2) + result[0]
	return result
}
