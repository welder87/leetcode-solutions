class Solution:
    def findUnsortedSubarray(self, nums: list[int]) -> int:
        mx = 0
        is_found = False
        for i in range(1, len(nums)):
            if nums[mx] > nums[i]:
                is_found = True
                break
            if nums[mx] < nums[i]:
                mx = i
        if not is_found:
            return 0
        mn = len(nums) - 1
        for i in range(len(nums) - 2, -1, -1):
            if nums[mn] < nums[i]:
                break
            if nums[mn] > nums[i]:
                mn = i
        return mn - mx + 1
