class Solution:
    def numSubarrayProductLessThanK(self, nums: list[int], k: int) -> int:
        if k <= 1:
            return 0
        ans = 0
        prod, left = 1, 0
        for right, num in enumerate(nums):
            prod *= num
            while prod >= k:
                prod //= nums[left]
                left += 1
            ans += right - left + 1
        return ans
