package problem56

import "sort"

// Time complexity: O(n log n). Space complexity: O(n log n).
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0, len(intervals))
	intermediate := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		start := max(intermediate[0], intervals[i][0])
		end := min(intermediate[1], intervals[i][1])
		if start <= end {
			intermediate[0] = min(intermediate[0], intervals[i][0])
			intermediate[1] = max(intermediate[1], intervals[i][1])
		} else {
			ans = append(ans, intermediate)
			intermediate = []int{intervals[i][0], intervals[i][1]}
		}
	}
	ans = append(ans, intermediate)
	return ans
}

// Time complexity: O(n log n). Space complexity: O(n log n).
// Solution: https://leetcode.com/problems/merge-intervals/editorial/
func mergeV1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}
