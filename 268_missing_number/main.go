package problem268

func missingNumber(nums []int) int {
	minNum := -1
	var maxNum, sum int
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if minNum == -1 || num < minNum {
			minNum = num
		}
		sum += num
	}
	fullSum := 0
	for i := minNum; i <= maxNum; i++ {
		fullSum += i
	}
	if fullSum == sum {
		if minNum != 0 {
			return 0
		}
		return maxNum + 1
	}
	return fullSum - sum
}
