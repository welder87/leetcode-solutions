package problem1991

func findMiddleIndex(nums []int) int {
	prefixSum := make([]int, len(nums)+1)
	for idx, num := range nums {
		prefixSum[idx+1] = prefixSum[idx] + num
	}
	midIdx := -1
	for idx := range nums {
		sum1 := prefixSum[idx]-prefixSum[0]
		sum2 := prefixSum[len(prefixSum)-1]-prefixSum[idx+1]
		if sum1 == sum2 {
			midIdx = idx
			break
		}
	}
	return midIdx
}
