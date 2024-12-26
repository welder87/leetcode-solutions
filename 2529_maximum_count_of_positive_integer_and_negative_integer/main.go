package problem2529

func maximumCount(nums []int) int {
	if nums[0] > 0 || nums[0] < 0 && nums[len(nums)-1] < 0 {
		return len(nums)
	}
	left, right, zeroStart, zeroEnd := 0, len(nums)-1, -1, -1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < 0 && nums[mid+1] > 0 {
			return max(len(nums)-1-mid, mid+1)
		} else if nums[mid] == 0 {
			zeroStart = mid
			right = mid - 1
		} else if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	left, right = zeroStart, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == 0 {
			zeroEnd = mid
			left = mid + 1
		} else if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if zeroStart >= 0 && zeroEnd >= 0 {
		return max(len(nums)-1-zeroEnd, zeroStart)
	} else if zeroStart >= 0 && zeroEnd < 0 {
		return max(len(nums)-1-zeroEnd, zeroStart)
	}
	return len(nums)
}
