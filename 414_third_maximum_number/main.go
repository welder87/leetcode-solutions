package problem414

func thirdMax(nums []int) int {
	maxs := make([]int, 0, 3)
	for _, num := range nums {
		if len(maxs) == 0 {
			maxs = append(maxs, num)
		} else if len(maxs) == 1 {
			if maxs[0] < num {
				maxs = append(maxs, maxs[0])
				maxs[0] = num
			} else if maxs[0] > num {
				maxs = append(maxs, num)
			}
		} else if len(maxs) == 2 {
			if maxs[0] < num {
				maxs = append(maxs, maxs[1])
				maxs[1] = maxs[0]
				maxs[0] = num
			} else if maxs[1] < num {
				maxs = append(maxs, maxs[1])
				maxs[1] = num
			} else if maxs[0] > num && maxs[1] > num {
				maxs = append(maxs, num)
			}
		} else {
			if maxs[0] < num {
				maxs[2] = maxs[1]
				maxs[1] = maxs[0]
				maxs[0] = num
			} else if maxs[1] < num {
				maxs[2] = maxs[1]
				maxs[1] = num
			} else if maxs[2] < num {
				maxs[2] = num
			}
		}
	}
	if len(maxs) < 3 {
		return maxs[0]
	}
	return maxs[2]
}
