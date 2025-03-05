class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        # Time complexity: O(n + m log m + n). Space complexity: O(m + k).
        unique_nums = sorted(set(nums))
        ans = len(unique_nums)
        for idx, num in enumerate(unique_nums):
            nums[idx] = num
        for idx in range(ans, len(nums)):
            nums[idx] = None
        return ans

    def removeDuplicatesV1(self, nums: list[int]) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        length = len(nums)
        j = 1
        for i in range(length):
            while j < length and nums[i] == nums[j]:
                j += 1
            if j < length:
                nums[i + 1] = nums[j]
            else:
                break
        ans = i + 1
        for i in range(ans, length):
            nums[i] = None
        return ans
