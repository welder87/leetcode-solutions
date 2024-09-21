package countpairs

import (
	"sort"
)

func CountPairs(nums []int, target int) int {
	sort.Ints(nums)
	var counter int
	for index, num := range nums {
		if num == target {
			counter = index + 1
			break
		}
	}
	return counter * (counter - 1) / 2
}
