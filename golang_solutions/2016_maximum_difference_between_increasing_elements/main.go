package problem2016

func maximumDifference(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	minDiff := nums[0]
	maxDiff := -1
	for i := 1; i < len(nums); i++ {
		if diff := nums[i] - minDiff; diff > maxDiff {
			maxDiff = diff
		} else if minDiff > nums[i] {
			minDiff = nums[i]
		}
	}
	if maxDiff <= 0 {
		maxDiff = -1
	}
	return maxDiff
}
