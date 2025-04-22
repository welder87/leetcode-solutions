package problem167

// Time complexity: O(n log n). Space complexity: O(1).
// modified itmo
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		num := target - numbers[i]
		left, right := -1, len(numbers)
		for right > left+1 {
			mid := left + (right-left)>>1
			if numbers[mid] == num && mid != i {
				return []int{i + 1, mid + 1}
			}
			if numbers[mid] <= num {
				left = mid
			} else {
				right = mid
			}
		}
	}
	return []int{}
}

// Time complexity: O(n log n). Space complexity: O(1).
// modified leetcode
func twoSumV1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		num := target - numbers[i]
		left, right := 0, len(numbers)-1
		for left < right {
			mid := left + (right-left)>>1
			if numbers[mid] == num && mid != i {
				return []int{min(mid, i) + 1, max(mid, i) + 1}
			}
			if numbers[mid] <= num {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return []int{}
}

// Time complexity: O(n). Space complexity: O(1).
func twoSumV2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sm := numbers[i] + numbers[j]
		if sm == target {
			return []int{i + 1, j + 1}
		}
		if sm < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
