package minimumabsdifference

import "sort"

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	return collectPairs(arr, findMinDistance(arr))
}

func MinimumAbsDifference1(arr []int) [][]int {
	new_arr := quickSort(arr)
	return collectPairs(new_arr, findMinDistance(new_arr))
}

func collectPairs(arr []int, minDistance int) [][]int {
	var result [][]int
	for i := 0; i < len(arr); i++ {
		j := i + 1
		if j >= len(arr) {
			break
		}
		if distance := absValue(arr[i], arr[j]); distance == minDistance {
			result = append(result, []int{arr[i], arr[j]})
		}
	}
	return result
}

func absValue(a int, b int) int {
	if val := a - b; val >= 0 {
		return val
	} else {
		return -val
	}
}

func findMinDistance(arr []int) int {
	isStart := true
	minDistance := 0
	for i := 0; i < len(arr); i++ {
		j := i + 1
		if j >= len(arr) {
			break
		}
		if distance := absValue(arr[i], arr[j]); isStart || distance < minDistance {
			minDistance = distance
			isStart = false
		}
	}
	return minDistance
}

func quickSort(arr []int) []int {
	lengthOfVector := len(arr)
	if lengthOfVector <= 1 {
		return arr
	}
	middleIndex := lengthOfVector / 2
	pivot := arr[middleIndex]
	new := make([]int, 0, lengthOfVector)
	left := make([]int, 0, lengthOfVector)
	middle := make([]int, 0, lengthOfVector)
	right := make([]int, 0, lengthOfVector)

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			middle = append(middle, val)
		} else {
			right = append(right, val)
		}
	}
	new = append(new, quickSort(left)...)
	new = append(new, middle...)
	new = append(new, quickSort(right)...)
	return new
}
