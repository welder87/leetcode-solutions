class Solution:
    def shuffle(self, nums: list[int], n: int) -> list[int]:
        # Time complexity: O(n). Space complexity: O(n).
        res = [None] * len(nums)
        i, j = 0, n
        for k in range(len(nums)):
            if k % 2 == 0:
                res[k] = nums[i]
                i += 1
            else:
                res[k] = nums[j]
                j += 1
        return res
