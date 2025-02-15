package problem713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans := 0
	prod, left := 1, 0
	for right, num := range nums {
		prod *= num
		for prod >= k {
			prod /= nums[left]
			left += 1
		}
		ans += right - left + 1
	}
	return ans
}

func numSubarrayProductLessThanKV1(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans := 0
	prod, left := 1, 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for prod >= k {
			prod /= nums[left]
			left += 1
		}
		ans += right - left + 1
	}
	return ans
}
