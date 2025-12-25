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

    def mergeV1(self, intervals: list[list[int]]) -> list[list[int]]:
        # Time complexity: O(n log n). Space complexity: O(n).
        # Solution: https://leetcode.com/problems/merge-intervals/editorial/
        intervals.sort(key=lambda x: x[0])
        merged = []
        for interval in intervals:
            # if the list of merged intervals is empty or if the current
            # interval does not overlap with the previous, simply append it.
            if not merged or merged[-1][1] < interval[0]:
                merged.append(interval)
            else:
                # otherwise, there is overlap, so we merge the current and previous
                # intervals.
                merged[-1][1] = max(merged[-1][1], interval[1])
        return merged
