from collections import Counter


class Solution:
    def singleNumber(self, nums: list[int]) -> int:
        # Time complexity: O(n + m). Space complexity: O(m).
        for num, cnt in Counter(nums).items():
            if cnt == 1:
                return num
        return 0

    def singleNumberV1(self, nums: list[int]) -> int:
        # Time complexity: O(n^2). Space complexity: O(1).
        for num1 in nums:
            cnt = 0
            for num2 in nums:
                if num1 == num2:
                    cnt += 1
            if cnt == 1:
                return num1
        return 0

    def singleNumberV2(self, nums: list[int]) -> int:
        # Time complexity: O(n log n + n ). Space complexity: O(n log n).
        if len(nums) == 1:
            return nums[0]
        nums.sort()
        i, j = 0, 1
        while j < len(nums):
            if nums[i] != nums[j]:
                return nums[i]
            i += 2
            j += 2
        return nums[i]

    def singleNumberV3(self, nums: list[int]) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        res = 0
        for num in nums:
            res ^= num
        return res
