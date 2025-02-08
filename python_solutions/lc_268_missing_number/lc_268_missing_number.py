class Solution:
    def missingNumber(self, nums: list[int]) -> int:
        # Time complexity: O(n^2). Space complexity: O(1).
        for num in range(len(nums) + 1):
            if num not in nums:
                return num
        return -1

    def missingNumberV1(self, nums: list[int]) -> int:
        # Time complexity: O(2n). Space complexity: O(n).
        unique_nums = set(nums)
        for num in range(len(nums) + 1):
            if num not in unique_nums:
                return num
        return -1

    def missingNumberV2(self, nums: list[int]) -> int:
        # Time complexity: O(n log n + n + 1). Space complexity: O(n log n).
        length = len(nums)
        nums.sort()
        for num in range(len(nums) + 1):
            if length == num or num != nums[num]:
                return num
        return -1

    def missingNumberV3(self, nums: list[int]) -> int:
        # Time complexity: O(2n). Space complexity: O(1).
        current = sum(nums)
        reference = sum(num for num in range(len(nums) + 1))
        return abs(reference - current)

    def missingNumberV4(self, nums: list[int]) -> int:
        # Time complexity: O(2n). Space complexity: O(1).
        current = 0
        for num in nums:
            current ^= num
        reference = 0
        for num in range(len(nums) + 1):
            reference ^= num
        return current ^ reference
