package problem2485

func pivotInteger(n int) int {
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + i
	}

	for i := 1; i <= n; i++ {
		sum1 := prefixSum[i] - prefixSum[0]
		sum2 := prefixSum[n] - prefixSum[i-1]
		if sum1 == sum2 {
			return i
		}
	}
	return -1
}

func pivotIntegerV1(n int) int {
	sum := n * (n + 1) >> 1
	for i := 1; i <= n; i++ {
		sum2 := i * (i + 1) >> 1
		sum1 := sum - (i - 1) * i >> 1
		if sum1 == sum2 {
			return i
		}
	}
	return -1
}
