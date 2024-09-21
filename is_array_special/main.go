package isarrayspecial

func IsArraySpecial(nums []int) bool {
	var next int
	var first int
	var second int
	lengthOfNums := len(nums)
	for i := 0; i < lengthOfNums; i++ {
		if next = i + 1; next >= lengthOfNums {
			break
		}
		first = nums[i]%2
		second = nums[next]%2
		if (first == 0 && second == 0) || (first != 0 && second != 0) {
			return false
		}
	}
	return true
}
