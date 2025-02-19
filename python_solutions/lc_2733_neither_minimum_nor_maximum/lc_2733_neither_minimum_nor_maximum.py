class Solution:
    def findNonMinOrMax(self, nums: list[int]) -> int:
        # Time complexity: O(n log n). Space complexity: O(n).
        if len(nums) < 3:
            return -1
        nums.sort()
        return nums[1]

    def findNonMinOrMaxV1(self, nums: list[int]) -> int:
        # Time complexity: O(2n). Space complexity: O(m).
        if len(nums) < 3:
            return -1
        sorted_nums = [False] * (max(nums) + 1)
        for num in nums:
            sorted_nums[num] = True
        counter = 0
        for num, is_exist in enumerate(sorted_nums):
            if is_exist:
                counter += 1
            if counter == 2:
                return num
        return -1

    def findNonMinOrMaxV2(self, nums: list[int]) -> int:
        # Time complexity: O(2n). Space complexity: O(1).
        min_, max_ = nums[0], nums[0]
        for num in nums:
            if min_ > num:
                min_ = num
            if max_ < num:
                max_ = num
        for num in nums:
            if min_ < num < max_:
                return num
        return -1
