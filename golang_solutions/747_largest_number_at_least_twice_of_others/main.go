package problem747

func dominantIndex(nums []int) int {
	first, second := 0, 1
	if nums[0] < nums[1] {
		first, second = 1, 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[first] {
			second = first
			first = i
		} else if nums[i] > nums[second] {
			second = i
		}
	}
	if nums[second] == 0 {
		return first
	}
	if nums[first]/nums[second] >= 2 {
		return first
	}
	return -1
}
