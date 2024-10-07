package removeelement

func RemoveElement(nums []int, val int) int {
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
