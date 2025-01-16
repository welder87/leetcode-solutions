package problem643

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	i := 0
	j := k - 1
	for n := i; n <= j; n++ {
		sum += nums[n]
	}
	maxSum := sum
	sum -= nums[i]
	i++
	j++
	for j < len(nums) {
		sum += nums[j]
		maxSum = max(maxSum, sum)
		sum -= nums[i]
		i++
		j++
	}
	return float64(maxSum) / float64(k)
}
