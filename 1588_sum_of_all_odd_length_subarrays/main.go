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
