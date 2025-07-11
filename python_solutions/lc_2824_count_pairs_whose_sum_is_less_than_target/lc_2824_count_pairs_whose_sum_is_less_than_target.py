class Solution:
    def countPairs(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n*n). Space complexity: O(1).
        if len(nums) <= 1:
            return 0
        ans = 0
        for i in range(0, len(nums) - 1):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] < target:
                    ans += 1
        return ans
