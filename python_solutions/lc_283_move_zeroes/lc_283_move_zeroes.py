class Solution:
    def moveZeroes(self, nums: list[int]) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        if len(nums) <= 1:
            return
        i, j = 0, 1
        while i < len(nums) and j < len(nums):
            if nums[i] == nums[j] == 0:
                j += 1
            elif nums[i] == 0 and nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            else:
                i += 1
                j += 1

    def moveZeroesV1(self, nums: list[int]) -> None:
        # Time complexity: O(2n). Space complexity: O(n).
        non_zero_nums = [None] * len(nums)
        i = 0
        for num in nums:
            if num != 0:
                non_zero_nums[i] = num
                i += 1
        for i, num in enumerate(non_zero_nums):
            nums[i] = 0 if num is None else num
