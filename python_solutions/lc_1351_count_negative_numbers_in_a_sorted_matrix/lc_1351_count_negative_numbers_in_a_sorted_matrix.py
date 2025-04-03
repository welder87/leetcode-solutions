class Solution:
    def countNegatives(self, grid: list[list[int]]) -> int:
        # Time complexity: O(m log n). Space complexity: O(1).
        # modified leetcode
        ans = 0
        for nums in grid:
            length = len(nums)
            left, right = 0, length - 1
            while left < right:
                mid = left + (right - left) // 2
                if nums[mid] >= 0:
                    left = mid + 1
                else:
                    right = mid
            if nums[left] < 0:
                ans += length - left
        return ans
