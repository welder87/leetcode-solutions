package problem697

// Time complexity: O(n + m). Space complexity: O(m).
func findShortestSubArray(nums []int) int {
	mapItems1 := make(map[int]int, len(nums))
	mapItems2 := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mapItems1[nums[i]]++
		if _, ok := mapItems2[nums[i]]; !ok {
			mapItems2[nums[i]]= []int{i, -1}
		} else {
			mapItems2[nums[i]][1] = i
		}
	}
	maxCnt, degree := 0, 0
	for num, cnt := range mapItems1 {
		idxs := mapItems2[num]
		cur := 1
		if idxs[1] != -1 {
			cur = idxs[1] - idxs[0] + 1
		}
		if maxCnt == 0 {
			maxCnt = cnt
			degree = cur
		} else if maxCnt == cnt {
			if degree > cur {
				degree = cur
			}
		} else if maxCnt < cnt {
			maxCnt = cnt
			degree = cur
		}
	}
	return degree
}
