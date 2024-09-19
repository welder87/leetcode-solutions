package intersectiontwoarrays2

func Intersect(nums1 []int, nums2 []int) []int {
	numsWithCount1 := make(map[int]int)
	for _, num := range nums1 {
		if _, ok := numsWithCount1[num]; ok {
			numsWithCount1[num]++
		} else {
			numsWithCount1[num] = 1
		}
	}
	vectorLength := len(nums1)
	if len(nums1) > len(nums2) {
		vectorLength = len(nums2)
	}
	result := make([]int, 0, vectorLength)
	for _, num := range nums2 {
		if count, ok := numsWithCount1[num]; ok {
			if count <= 1 {
				delete(numsWithCount1, num)
			} else {
				numsWithCount1[num]--
			}
			result = append(result, num)
		}
	}
	return result
}

func Intersect1(nums1 []int, nums2 []int) []int {
	numsWithCount1 := make(map[int]int)
	for _, val := range nums1 {
		if _, ok := numsWithCount1[val]; ok {
			numsWithCount1[val]++
		} else {
			numsWithCount1[val] = 1
		}
	}
	numsWithCount2 := make(map[int]int)
	for _, val := range nums2 {
		if _, ok := numsWithCount2[val]; ok {
			numsWithCount2[val]++
		} else {
			numsWithCount2[val] = 1
		}
	}
	vectorLength := len(nums1)
	if len(nums1) > len(nums2) {
		vectorLength = len(nums2)
	}
	result := make([]int, 0, vectorLength)
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
