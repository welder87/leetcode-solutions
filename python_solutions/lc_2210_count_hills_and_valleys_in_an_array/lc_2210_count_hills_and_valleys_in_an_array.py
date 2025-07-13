class Solution:
    def countHillValley(self, nums: list[int]) -> int:
        counter = 0
        left, mid = 0, 1
        for right in range(2, len(nums)):
            while (
                left < len(nums)
                and (nums[left] <= nums[mid] <= nums[right])
                or (nums[left] >= nums[mid] >= nums[right])
            ):
                left += 1
                mid += 1
            counter += 1
        return counter
