package problem1394

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
