class Solution:
    def thirdMax(self, nums: list[int]) -> int:
        # Time complexity: O(n + m log m). Space complexity: O(m log m + m).
        distinct_nums = sorted(set(nums), reverse=True)
        if len(distinct_nums) < 3:
            return distinct_nums[0]
        return distinct_nums[2]

    def thirdMaxV1(self, nums: list[int]) -> int:
        # Time complexity: O(n log n + n). Space complexity: O(n log n + n).
        if len(nums) < 3:
            return max(nums)
        nums.sort()
        # -2, -2, -2, 1, 4
        third, second, first = None, None, nums[-1]
        for num in reversed(nums):
            if second is None and num < first:
                second = num
            elif second is not None and num < second and third is None:
                third = num
                break
        if third is not None:
            return third
        return first
