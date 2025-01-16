package problem1588

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 1; i < len(arr)+1; i = i + 2 {
		for j := 0; j < len(arr)-i+1; j++ {
			for k := j; k < j+i; k++ {
				sum += arr[k]
			}
		}
	}
	return sum
}

func sumOddLengthSubarraysV1(arr []int) int {
	prefixSum := make([]int, len(arr)+1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = arr[i-1] + prefixSum[i-1]
	}
	sum := 0
	for i := 1; i < len(arr)+1; i = i + 2 {
		for j, k := 0, i; k < len(prefixSum); k, j = k+1, j+1 {
			sum += prefixSum[k] - prefixSum[j]
		}
	}
	return sum
}
