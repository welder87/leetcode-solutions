package problem57

// Time complexity: O(n). Space complexity: O(n).
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0, len(intervals))
	i := 0
	// этап до того, как новый интервал может быть добавлен
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}
	// этап поиска пересечений нового интервала с другими и слияния, соответственно
	ans = append(ans, newInterval)
	for i < len(intervals) &&
		max(newInterval[0], intervals[i][0]) <= min(newInterval[1], intervals[i][1]) {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	// этап после того, как новый интервал добавлен
	for i < len(intervals) {
		ans = append(ans, intervals[i])
		i++
	}
	return ans
}
