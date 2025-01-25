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

    def findMaxKV1(self, nums: list[int]) -> int:
        pos_nums = {num for num in nums if num > 0}
        if not pos_nums:
            return -1
        maxs = -1
        for num in nums:
            if num < 0 and (val := abs(num)) in pos_nums and maxs < val:
                maxs = val
        return maxs

    def findMaxKV2(self, nums: list[int]) -> int:
        nums.sort()
        i, j = 0, len(nums) - 1
        while nums[i] < 0 and nums[j] > 0:
            val1, val2 = -nums[i], nums[j]
            if val1 == val2:
                return val2
            elif val1 > val2:
                i += 1
            else:
                j -= 1
        return -1
