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
