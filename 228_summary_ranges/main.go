package problem228

import (
	"strconv"
	"strings"
)

const pairLength = 2

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	summary := make([]string, 0, len(nums))
	pair := make([]int, 0, pairLength)
	for _, num := range nums {
		if len(pair) == 0 {
			pair = append(pair, num)
			continue
		}
		if len(pair) == 1 && num-pair[0] == 1 {
			pair = append(pair, num)
			continue
		}
		if len(pair) == 2 && num-pair[1] == 1 {
			pair[1] = num
			continue
		}
		summary = append(summary, makeStringFromArrayOfInt(&pair))
		pair = make([]int, 0, pairLength)
		pair = append(pair, num)
	}
	if len(pair) > 0 {
		summary = append(summary, makeStringFromArrayOfInt(&pair))
	}
	return summary
}

func makeStringFromArrayOfInt(pair *[]int) string {
	strAr := make([]string, 0, pairLength)
	for _, val := range *pair {
		strAr = append(strAr, strconv.Itoa(val))
	}
	return strings.Join(strAr, "->")
}
