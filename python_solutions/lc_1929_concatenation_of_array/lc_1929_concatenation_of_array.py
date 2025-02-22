class Solution:
    def getConcatenation(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(2n).
        length = len(nums)
        ans = [None] * (length * 2)
        for idx, num in enumerate(nums):
            ans[idx], ans[idx + length] = num, num
        return ans

    def getConcatenationV1(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(2n).
        length = len(nums) * 2
        i = 0
        while length > len(nums):
            nums.append(nums[i])
            i += 1
        return nums

    def getConcatenationV2(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(2n).
        return nums * 2
