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
