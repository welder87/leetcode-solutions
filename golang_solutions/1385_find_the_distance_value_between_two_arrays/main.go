package problem1385

// Time complexity: O(n * m). Space complexity: O(1).
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0
	for i := 0; i < len(arr1); i++ {
		isSuited := true
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				isSuited = false
				break
			}
		}
		if isSuited {
			ans++
		}
	}
	return ans
}

func abs(res int) int {
	if res < 0 {
		return res * -1
	}
	return res
}
