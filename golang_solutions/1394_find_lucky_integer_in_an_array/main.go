package problem1394

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
