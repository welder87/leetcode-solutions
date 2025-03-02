package problem1394

import "sort"

// Time complexity: O(n + m). Space complexity: O(m).
func findLucky(arr []int) int {
	counter := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		counter[arr[i]]++
	}
	ans := -1
	for num, cnt := range counter {
		if num == cnt && num > ans {
			ans = num
		}
	}
	return ans
}

// Time complexity: O(2n + m). Space complexity: O(m).
func findLuckyV1(arr []int) int {
	maxNum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}
	counter := make([]int, maxNum+1)
	for i := 0; i < len(arr); i++ {
		counter[arr[i]]++
	}
	ans := -1
	for num := 1; num < len(counter); num++ {
		if num == counter[num] && num > ans {
			ans = num
		}
	}
	return ans
}

// Time complexity: O(n log n + n). Space complexity: O(n).
func findLuckyV2(arr []int) int {
	if len(arr) == 1 {
		if arr[0] == 1 {
			return 1
		}
		return -1
	}
	sort.Ints(arr)
	i := len(arr) - 1
	largest_lucky := -1
	for i >= 0 {
		counter, num := 0, arr[i]
		for i >= 0 && num == arr[i] {
			counter += 1
			i -= 1
		}
		if num == counter && largest_lucky < num {
			largest_lucky = num
		}
	}
	return largest_lucky
}
