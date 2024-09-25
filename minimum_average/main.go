package minimumaverage

import (
	"slices"
	"sort"
)

func MinimumAverage(nums []int) float64 {
	sort.Ints(nums)
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
