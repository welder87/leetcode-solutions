class Solution:
    def isMonotonic(self, nums: list[int]) -> bool:
        if len(nums) == 0:
            return False
        is_increasing, is_decreasing = True, True
        for i in range(1, len(nums)):
            if is_increasing and nums[i] > nums[i - 1]:
                is_increasing = False
            if is_decreasing and nums[i] < nums[i - 1]:
                is_decreasing = False
        return is_increasing or is_decreasing
