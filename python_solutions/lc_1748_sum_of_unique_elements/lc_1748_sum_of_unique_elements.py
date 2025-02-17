from collections import Counter


class Solution:
    def sumOfUniqueV1(self, nums: list[int]) -> int:
        sum_ = 0
        for num, cnt in Counter(nums).items():
            if cnt == 1:
                sum_ += num
        return sum_
