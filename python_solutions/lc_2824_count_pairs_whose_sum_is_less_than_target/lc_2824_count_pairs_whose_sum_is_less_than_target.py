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

    def countPairsV1(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n log n) + O(n log n). Space complexity: O(n).
        # modified itmo
        # [-7, -6, -2, -1, 2, 3, 5]
        if len(nums) <= 1:
            return 0
        nums.sort()
        ans = 0
        for i in range(len(nums) - 1):
            left, right = i + 1, len(nums)
            while left + 1 < right:
                mid = left + (right - left) // 2
                if i < mid and nums[i] + nums[mid] < target:
                    left = mid
                else:
                    right = mid
            if i < left and nums[i] + nums[left] < target:
                ans += left - i
        return ans
