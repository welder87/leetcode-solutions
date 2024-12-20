package problem912

// Bubble sort (time complexity O(n^2), space complexity O(1))
func sortArray(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// Counting sort (time complexity O(n), space complexity O(n))
func sortArrayV1(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	maxNum := nums[0]
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	positive := make([]int, absNum(maxNum)+1)
	negative := make([]int, absNum(minNum))
	for _, num := range nums {
		if num >= 0 {
			positive[num]++
		} else {
			negative[(num*-1)-1]++
		}
	}
	result := make([]int, 0, absNum(maxNum)+absNum(minNum)+1)
	for i := len(negative) - 1; i >= 0; i-- {
		if negative[i] > 0 {
			for j := 0; j < negative[i]; j++ {
				result = append(result, (i+1)*-1)
			}
		}
	}
	for i := 0; i < len(positive); i++ {
		if positive[i] > 0 {
			for j := 0; j < positive[i]; j++ {
				result = append(result, i)
			}
		}
	}
	return result
}

func absNum(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

// Quick sort (time complexity O(n log n), space complexity O(1))
func sortArrayV2(nums []int) []int {
	subSortArray(nums, 0, len(nums)-1)
	return nums
}

func subSortArray(nums []int, p int, r int) {
	if p < r {
		q := partition(nums, p, r)
		subSortArray(nums, p, q-1)
		subSortArray(nums, q+1, r)
	}
}

func partition(nums []int, p int, pivotIndex int) int {
	pivot := nums[pivotIndex]
	i := p - 1
	for j := p; j < pivotIndex; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[pivotIndex] = nums[pivotIndex], nums[i+1]
	return i + 1
}
