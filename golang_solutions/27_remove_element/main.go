package problem27

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] != val && nums[j] != val {
			i++
		} else if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			nums[j] = -1
			i++
			j--
		} else if nums[i] != val && nums[j] == val {
			nums[j] = -1
			i++
			j--
		} else {
			nums[j] = -1
			j--
		}
	}
	return j + 1
}

func removeElementV1(nums []int, val int) int {
	i, j, cnt := 0, len(nums)-1, 0
	for i <= j {
		if nums[j] == val {
			nums[j], j, cnt = -1, j-1, cnt+1
		} else if nums[i] == val {
			nums[i], nums[j], i, j, cnt = nums[j], -1, i+1, j-1, cnt+1
		} else {
			i++
		}
	}
	return len(nums) - cnt
}
