package problem88

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for i >= 0 || j >= 0 {
		if i < 0 || (j >= 0 && nums1[i] <= nums2[j]) {
			nums1[i+j+1] = nums2[j]
			j -= 1
		} else {
			nums1[i+j+1] = nums1[i]
			i -= 1
		}
	}
}

func mergeV1(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}
