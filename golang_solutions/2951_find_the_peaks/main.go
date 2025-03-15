package problem2951

// Time complexity: O(n). Space complexity: O(1).
func findPeaks(mountain []int) []int {
	ans := make([]int, 0, len(mountain))
	lastIndex := len(mountain) - 1
	for i := 1; i < lastIndex; i++ {
		if mountain[i-1] < mountain[i] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return ans
}
