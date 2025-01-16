package problem2553

func separateDigits(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		subAns := []int{}
		subNum := num / 10
		remainder := num % 10
		for subNum > 0 {
			subAns = append(subAns, remainder)
			remainder = subNum % 10
			subNum /= 10
		}
		if remainder > 0 {
			subAns = append(subAns, remainder)
		}
		for i := len(subAns) - 1; i >= 0; i-- {
			ans = append(ans, subAns[i])
		}
	}
	return ans
}
