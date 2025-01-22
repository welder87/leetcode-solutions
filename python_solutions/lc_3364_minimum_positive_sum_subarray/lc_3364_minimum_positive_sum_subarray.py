class Solution:
    def minimumSumSubarray(self, nums: list[int], l: int, r: int) -> int:
        prefix_sum = [0] * (len(nums) + 1)
        for i in range(len(nums)):
            prefix_sum[i + 1] = prefix_sum[i] + nums[i]
        min_sum = None
        for i in range(l, r + 1):
            for j in range(0, len(prefix_sum) - i):
                cur_sum = prefix_sum[j + i] - prefix_sum[j]
                if cur_sum > 0 and (min_sum is None or min_sum > cur_sum):
                    min_sum = cur_sum
        return -1 if min_sum is None else min_sum
