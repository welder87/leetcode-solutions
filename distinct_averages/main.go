package distinctaverages

func DistinctAverages(nums []int) int {
	nums = quickSort(nums)
	uniqueAverages := make(map[float64]struct{})
	i := 0
	j := len(nums) - 1
	for i < j {
		val := float64(nums[i]+nums[j]) / 2.0
		uniqueAverages[val] = struct{}{}
		i++
		j--
	}
	return len(uniqueAverages)
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[len(nums)/2]
	left := make([]int, 0, len(nums))
	right := make([]int, 0, len(nums))
	middle := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			right = append(right, nums[i])
		} else if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			middle = append(middle, nums[i])
		}
	}
	new := make([]int, 0, len(nums))
	new = append(new, quickSort(left)...)
	new = append(new, middle...)
	new = append(new, quickSort(right)...)
	return new
}
