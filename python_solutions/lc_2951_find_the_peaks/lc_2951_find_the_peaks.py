class Solution:
    def findPeaks(self, mountain: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(1).
        return [
            i
            for i in range(1, len(mountain) - 1)
            if mountain[i - 1] < mountain[i] > mountain[i + 1]
        ]
