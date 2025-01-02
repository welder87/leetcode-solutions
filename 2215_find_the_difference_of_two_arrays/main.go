package problem2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	uniqueNums1 := make(map[int]struct{}, len(nums1))
	for _, num := range nums1 {
		if _, ok := uniqueNums1[num]; !ok {
			uniqueNums1[num] = struct{}{}
		}
	}
	uniqueNums2 := make(map[int]struct{}, len(nums2))
	for _, num := range nums2 {
		if _, ok := uniqueNums2[num]; !ok {
			uniqueNums2[num] = struct{}{}
		}
	}
	result1 := make([]int, 0, len(nums1))
	result2 := make([]int, 0, len(nums2))
	for num := range uniqueNums1 {
		if _, ok := uniqueNums2[num]; !ok {
			result1 = append(result1, num)
		}
	}
	for num := range uniqueNums2 {
		if _, ok := uniqueNums1[num]; !ok {
			result2 = append(result2, num)
		}
	}
	result := make([][]int, 0, 2)
	result = append(result, result1)
	result = append(result, result2)
	return result
}
