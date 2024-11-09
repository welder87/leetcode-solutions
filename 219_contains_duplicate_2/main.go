package problem219

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	counter := make(map[int][]int)
	for idx, val := range nums {
		counter[val] = append(counter[val], idx)
	}
	for _, idxs := range counter {
		if len(idxs) <= 1 {
			continue
		}

		for i := 0; i < len(idxs); i++ {
			for j := i + 1; j < len(idxs); j++ {
				if (idxs[j] - idxs[i]) <= k {
					return true
				}
			}
		}
	}
	return false
}
