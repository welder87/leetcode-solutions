class Solution:
    def numSubarrayProductLessThanK(self, nums: list[int], k: int) -> int:
        if k == 0:
            return 0
        ans, prod, left = 0, 1, 0
        for right, num in enumerate(nums):
            prod *= num
            while prod >= k and left < len(nums):
                prod //= nums[left]
                if prod < 1:
                    prod = 1
                left += 1
            if left >= len(nums):
                break
            ans += right - left + 1
        return ans
