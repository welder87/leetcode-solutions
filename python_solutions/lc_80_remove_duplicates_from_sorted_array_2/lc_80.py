class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        # Time complexity: O(n + n). Space complexity: O(1).
        i, j = 0, 1
        cnt = 1
        counter = 0
        # first pass through the array
        # marking as empty value
        while i < len(nums) or j < len(nums):
            while j < len(nums) and nums[i] == nums[j]:
                cnt += 1
                if cnt > 2:
                    nums[j] = None
                    counter += 1
                j += 1
            i = j
            j += 1
            cnt = 1
        j = 1
        # second pass through the array
        # moving empty values
        for i in range(len(nums)):
            if nums[i] is not None:
                j += 1
                continue
            while j < len(nums) and nums[j] is None:
                j += 1
            if j < len(nums):
                nums[i], nums[j] = nums[j], nums[i]
        return len(nums) - counter
