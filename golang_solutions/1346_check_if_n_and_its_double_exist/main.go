package problem1346

import "sort"

// Time complexity: O(n log n + n log n). Space complexity: O(n).
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		left, right := -1, len(arr)
		for right > left + 1 {
			mid := left + (right - left) >> 1
			if val := 2 * arr[mid]; i != mid && arr[i] == val {
				return true
			} else if val < arr[i] {
				left = mid
			} else {
				right = mid
			}
		}
	}
	return false
}

