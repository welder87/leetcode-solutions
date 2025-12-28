package problem1288

import (
	"sort"
)

// Time complexity: O(n log n). Space complexity: O(n log n).
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	counter := 0
	pre := -1
	for i := range intervals {
		if intervals[i][1] > pre {
			counter++
			pre = intervals[i][1]
		}
	}
	return counter
}
