class Solution:
    def applyOperations(self, nums: list[int]) -> list[int]:
        # Time complexity: O(2n). Space complexity: O(1).
        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                nums[i], nums[i + 1] = nums[i] * 2, 0
        i, j = 0, 1
        while j < len(nums):
            if nums[i] == 0:
                if nums[j] != 0:
                    nums[i], nums[j] = nums[j], nums[i]
                    i += 1
            else:
                i += 1
            j += 1
        return nums

    def applyOperationsV1(self, nums: list[int]) -> list[int]:
        # Time complexity: O(2n). Space complexity: O(n).
        for i in range(len(nums) - 1):
            if nums[i] == nums[i + 1]:
                nums[i], nums[i + 1] = nums[i] * 2, 0
        ans = [0] * len(nums)
        idx = 0
        for num in nums:
            if num != 0:
                ans[idx] = num
                idx += 1
        return ans
