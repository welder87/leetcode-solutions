class Solution:
    def sortColors(self, nums: list[int]) -> None:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] > nums[j]:
                    nums[i], nums[j] = nums[j], nums[i]

    def sortColorsV1(self, nums: list[int]) -> None:
        length = len(nums)
        for i in range(length):
            has_swapped = False
            for j in range(0, length - i - 1):
                if nums[j] > nums[j + 1]:
                    nums[j], nums[j + 1] = nums[j + 1], nums[j]
                    has_swapped = True
            if not has_swapped:
                break
