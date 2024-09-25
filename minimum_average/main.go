package minimumaverage

import (
	"slices"
)

func MinimumAverage(nums []int) float64 {
	nums = quickSort(nums)
	average := make([]float64, 0, len(nums))
	i := 0
	j := len(nums) - 1
	for i < j {
		average = append(average, (float64(nums[i])+float64(nums[j]))/2)
		i++
		j--
	}
	return slices.Min(average)
}

func quickSort(arr []int) []int {
	lengthOfVector := len(arr)
	if lengthOfVector <= 1 {
		return arr
	}
	middleIndex := lengthOfVector / 2
	pivot := arr[middleIndex]
	new := make([]int, 0, lengthOfVector)
	left := make([]int, 0, lengthOfVector)
	middle := make([]int, 0, lengthOfVector)
	right := make([]int, 0, lengthOfVector)

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			middle = append(middle, val)
		} else {
			right = append(right, val)
		}
	}
	new = append(new, quickSort(left)...)
	new = append(new, middle...)
	new = append(new, quickSort(right)...)
	return new
}
