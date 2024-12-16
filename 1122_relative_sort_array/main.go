package problem1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := make(map[int]int, len(arr1))
	for _, num := range arr1 {
		counter[num]++
	}
	result := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		if cnt, ok := counter[num]; ok {
			for i := 0; i < cnt; i++ {
				result = append(result, num)
			}
			delete(counter, num)
		}
	}
	subResult := make([]int, 0, len(arr1))
	for num, cnt := range counter {
		for i := 0; i < cnt; i++ {
			subResult = append(subResult, num)
		}
		delete(counter, num)
	}
	sort.Ints(subResult)
	result = append(result, subResult...)
	return result
}
