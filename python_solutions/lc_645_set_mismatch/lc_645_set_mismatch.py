from collections import Counter


class Solution:
    def findErrorNums(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n + m). Space complexity: O(m).
        counter = Counter(nums)
        reference, repeat = None, None
        for num in range(1, len(nums) + 1):
            if (cnt := counter[num]) > 1:
                repeat = num
            elif cnt == 0:
                reference = num
        return [repeat, reference]

    def findErrorNumsV1(self, nums: list[int]) -> list[int]:
        # Time complexity: O(2n). Space complexity: O(n+1).
        counter = [0] * (len(nums) + 1)
        reference, repeat = None, None
        for num in nums:
            counter[num] += 1
        for num in range(1, len(counter)):
            if (cnt := counter[num]) > 1:
                repeat = num
            elif cnt == 0:
                reference = num
        return [repeat, reference]

    def findErrorNumsV2(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n log n + 2 n + 1). Space complexity: O(n log n).
        nums.sort()
        reference, repeat = None, None
        sum_with_repeat, reference_sum = sum(nums), 0
        for idx in range(1, len(nums) + 1):
            if repeat is None and nums[idx] == nums[idx - 1]:
                repeat = nums[idx]
            reference_sum += idx
        reference = reference_sum - (sum_with_repeat - repeat)
        return [repeat, reference]
