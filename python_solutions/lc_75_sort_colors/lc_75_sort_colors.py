class Solution:
    def sortColors(self, nums: list[int]) -> None:
        # Time complexity: O(n^2). Space complexity: O(1).
        # Bubble Sort
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] > nums[j]:
                    nums[i], nums[j] = nums[j], nums[i]

    def sortColorsV1(self, nums: list[int]) -> None:
        # Time complexity: O(n^2). Space complexity: O(1).
        # modified Bubble Sort
        length = len(nums)
        for i in range(length):
            has_swapped = False
            for j in range(0, length - i - 1):
                if nums[j] > nums[j + 1]:
                    nums[j], nums[j + 1] = nums[j + 1], nums[j]
                    has_swapped = True
            if not has_swapped:
                break

    def sortColorsV2(self, nums: list[int]) -> None:
        # Time complexity: O(n) + O(n). Space complexity: O(1).
        # Counting Sort
        counter = [0, 0, 0]
        for num in nums:
            counter[num] += 1
        j = 0
        for num, cnt in enumerate(counter):
            for i in range(cnt):
                nums[j + i] = num
            j += cnt
