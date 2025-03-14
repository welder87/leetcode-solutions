class Solution:
    def findPeakElement(self, nums: list[int]) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        if len(nums) == 1:
            return 0
        if len(nums) == 2:
            return 0 if nums[0] > nums[1] else 1

        for i in range(-1, len(nums)):
            if (
                (i == -1 and nums[i + 1] > nums[i + 2])
                or (i + 2 == len(nums) and nums[i] < nums[i + 1])
                or (nums[i] < nums[i + 1] > nums[i + 2])
            ):
                return i + 1
        return -1
