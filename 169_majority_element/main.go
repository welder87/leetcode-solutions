package problem169

func majorityElement(nums []int) int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}
	var elem, maxCnt int
	for num, cnt := range counter {
		if cnt > maxCnt {
			maxCnt = cnt
			elem = num
		}
	}
	return elem
}

// The solution uses the Moore Voting Algorithm
func majorityElementV1(nums []int) int {
	counter := 0
	elem := 0
	for _, num := range nums {
		if counter == 0 {
			elem = num
			counter++
		} else if elem == num {
			counter++
		} else {
			counter--
		}
	}
	return elem
}
