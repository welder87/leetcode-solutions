class Solution:
    def findMaxK(self, nums: list[int]) -> int:
        maxs = -1
        for i in range(len(nums) - 1):
            for j in range(i + 1, len(nums)):
                if nums[i] < 0 and nums[j] < 0 or nums[i] > 0 and nums[j] > 0:
                    continue
                if (val := abs(nums[i])) == abs(nums[j]) and maxs < val:
                    maxs = val
        return maxs
