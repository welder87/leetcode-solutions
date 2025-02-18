class Solution:
    def thirdMax(self, nums: list[int]) -> int:
        # Time complexity: O(n + m log m). Space complexity: O(m log m + m).
        distinct_nums = sorted(set(nums), reverse=True)
        if len(distinct_nums) < 3:
            return distinct_nums[0]
        return distinct_nums[2]
