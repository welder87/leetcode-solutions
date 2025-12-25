class Solution:
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        # Time complexity: O(n log n). Space complexity: O(n).
        intervals.sort(key=lambda x: x[0])
        ans = []
        intermediate = intervals[0]
        for i in range(1, len(intervals)):
            start = max(intermediate[0], intervals[i][0])
            end = min(intermediate[1], intervals[i][1])
            if start <= end:
                intermediate[0] = min(intermediate[0], intervals[i][0])
                intermediate[1] = max(intermediate[1], intervals[i][1])
            else:
                ans.append(intermediate)
                intermediate = intervals[i]
        ans.append(intermediate)
        return ans
