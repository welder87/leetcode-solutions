package problem350

func intersect(nums1 []int, nums2 []int) []int {
	numsWithCount := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		numsWithCount[num]++
	}
	result := make([]int, 0, max(len(nums1), len(nums2)))
	for _, num := range nums2 {
		if count, ok := numsWithCount[num]; ok {
			if count <= 1 {
				delete(numsWithCount, num)
			} else {
				numsWithCount[num]--
			}
			result = append(result, num)
		}
	}
	return result
}

func intersectV1(nums1 []int, nums2 []int) []int {
	numsWithCount1 := make(map[int]int, len(nums1))
	for _, val := range nums1 {
		numsWithCount1[val]++
	}
	numsWithCount2 := make(map[int]int, len(nums2))
	for _, val := range nums2 {
		numsWithCount2[val]++
	}
	result := make([]int, 0, max(len(nums1), len(nums2)))
	for num1, count1 := range numsWithCount1 {
		if count2, ok := numsWithCount2[num1]; ok {
			count := count2
			if count2 > count1 {
				count = count1
			}
			for i := 0; i < count; i++ {
				result = append(result, num1)
			}
		}
	}
	return result
}
