package problem349

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]struct{}, len(nums1) + len(nums2))
	for _, num := range nums1 {
		hashMap[num] = struct{}{}
	}
	vector := make([]int, 0, max(len(nums2), len(nums2)))
	for _, num := range nums2 {
		if _, ok := hashMap[num]; ok {
			vector = append(vector, num)
			delete(hashMap, num)
		}
	}
	return vector
}

func intersectionV1(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]struct{})
	for _, num := range nums1 {
		if val, ok := map1[num]; !ok {
			map1[num] = val
		}
	}
	map2 := make(map[int]struct{})
	for _, num := range nums2 {
		if val, ok := map2[num]; !ok {
			map2[num] = val
		}
	}
	vectorLength := len(nums1)
	if len(nums1) > len(nums2) {
		vectorLength = len(nums2)
	}
	vector := make([]int, 0, vectorLength)
	for num := range map1 {
		if _, ok := map2[num]; ok {
			vector = append(vector, num)
		}
	}
	return vector
}
