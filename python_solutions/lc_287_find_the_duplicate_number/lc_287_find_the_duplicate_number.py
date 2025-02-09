from collections import Counter


class Solution:
    def findDuplicate(self, nums: list[int]) -> int:
        # Time complexity: O(n + m). Space complexity: O(m).
        for num, cnt in Counter(nums).items():
            if cnt > 1:
                return num
        return 0

    def findDuplicateV1(self, nums: list[int]) -> int:
        # Time complexity: O(n^2). Space complexity: O(1).
        for num1 in nums:
            cnt = 0
            for num2 in nums:
                if num1 == num2:
                    cnt += 1
            if cnt > 1:
                return num1
        return 0

    def findDuplicateV2(self, nums: list[int]) -> int:
        # Time complexity: O(n log n + n). Space complexity: O(n log n).
        nums.sort()
        length = len(nums)
        i, j = 0, 1
        while i < length and j < length:
            if nums[i] == nums[j]:
                return nums[i]
            i += 1
            j += 1
        return 0
