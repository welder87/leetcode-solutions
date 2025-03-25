class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2 log n). Space complexity: O(1).
        if len(nums) == 0:
            return [-1, -1]
        left, right = self.left_bs(nums, target), self.right_bs(nums, target)
        return [-1, -1] if left > right else [left, right]

    def right_bs(self, nums: list[int], target: int) -> int:
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] <= target:
                left = mid
            else:
                right = mid
        return left

    def left_bs(self, nums: list[int], target: int) -> int:
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid
            else:
                right = mid
        return right
