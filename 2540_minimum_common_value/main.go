package problem2540

func getCommon(nums1 []int, nums2 []int) int {
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}

func getCommonV1(nums1 []int, nums2 []int) int {
	n1 := &nums1
	n2 := &nums2
	if len(nums1) < len(nums2) {
		n1 = &nums2
		n2 = &nums1
	}
	uniqueNums := make(map[int]struct{}, len(*n1))
	for _, num := range *n1 {
		uniqueNums[num] = struct{}{}
	}
	for _, num := range *n2 {
		if _, ok := uniqueNums[num]; ok {
			return num
		}
	}
	return -1
}
