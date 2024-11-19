package problem448

// naive. solution with hash table
func findDisappearedNumbers(nums []int) []int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	ans := make([]int, 0, len(nums))
	for i := 1; i <= len(nums); i++ {
		if _, ok := set[i]; !ok {
			ans = append(ans, i)
		}
	}
	return ans
}
