package problem724

func pivotIndex(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0
	for idx := 0; idx < len(nums); idx++ {
		prefixSum[idx+1] = prefixSum[idx] + nums[idx]
	}
	lastIndex := len(prefixSum) - 1
	for idx := 1; idx < len(prefixSum); idx++ {
		if prefixSum[idx-1]-prefixSum[0] == prefixSum[lastIndex]-prefixSum[idx] {
			return idx - 1
		}
	}
	return -1
}

func pivotIndexV1(nums []int) int {
	leftSum := 0
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}
	for idx, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return idx
		}
		leftSum += num
	}
	return -1
}
