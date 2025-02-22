from collections import Counter


class Solution:
    def frequencySort(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n + m log m + n). Space complexity: O(2m).
        sorted_nums = sorted(Counter(nums).items(), key=lambda x: (x[1], -x[0]))
        i = 0
        for num, cnt in sorted_nums:
            for _ in range(cnt):
                nums[i] = num
                i += 1
        return nums
