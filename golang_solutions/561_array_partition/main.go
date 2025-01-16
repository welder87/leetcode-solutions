package problem561

func arrayPairSum(nums []int) int {
	maxNum := -1
	minNum := nums[0]
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
		if minNum > num {
			minNum = num
		}
	}
	sortedNums := make([]int, 0, len(nums))
	if minNum < 0 {
		negativeCounter := make([]int, minNum*-1+1)
		for _, num := range nums {
			if num < 0 {
				negativeCounter[num*-1]++
			}
		}
		i := len(negativeCounter) - 1
		for i > 0 {
			if cnt := negativeCounter[i]; cnt > 0 {
				sortedNums = append(sortedNums, i*-1)
				negativeCounter[i]--
			} else {
				i--
			}
		}
	}
	if maxNum >= 0 {
		positiveCounter := make([]int, maxNum+1)
		for _, num := range nums {
			if num >= 0 {
				positiveCounter[num]++
			}
		}
		i := 0
		for i < len(positiveCounter) {
			if cnt := positiveCounter[i]; cnt > 0 {
				sortedNums = append(sortedNums, i)
				positiveCounter[i]--
			} else {
				i++
			}
		}
	}
	maximizedSum := 0
	for i, j := 0, 1; i < len(sortedNums); i, j = i+2, j+2 {
		maximizedSum += min(sortedNums[i], sortedNums[j])
	}
	return maximizedSum
}
