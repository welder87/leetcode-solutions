class Solution:
    def rearrangeArray(self, nums: list[int]) -> list[int]:
        prev = None
        for i in range(0, len(nums), 2):
            j = i + 1
            if nums[i] > 0:
                continue
            while j < len(nums) and nums[j] < 0:
                j += 1
            if j < len(nums):
                nums[i], nums[j] = nums[j], nums[i]
        return nums

    def rearrangeArrayV1(self, nums: list[int]) -> list[int]:
        i, j = 0, 1
        for i in range(len(nums)):
            if i % 2 == 0:
                if nums[i] < 0:
                    while j < len(nums) and nums[j] < 0:
                        j += 1
                    if j < len(nums):
                        nums[i], nums[j] = nums[j], nums[i]
            else:
                if nums[i] > 0:
                    while j < len(nums) and nums[j] > 0:
                        j += 1
                    if j < len(nums):
                        nums[i], nums[j] = nums[j], nums[i]
            j += 1
        return nums
