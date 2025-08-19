class Solution:
    def sortArray(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n log n). Space complexity: O(3n).
        # Method: Recursive Merge Sort With Slices
        return self._merge_sort(nums)

    def _merge_sort(self, nums: list[int]):
        if len(nums) < 2:
            return nums
        mid = len(nums) // 2
        ar1, ar2 = nums[:mid], nums[mid:]
        return self._merge_two_arrays(self._merge_sort(ar1), self._merge_sort(ar2))

    def _merge_two_arrays(self, ar1: list[int], ar2: list[int]) -> list[int]:
        len1, len2 = len(ar1), len(ar2)
        i, j = 0, 0
        ans = [0] * (len1 + len2)
        while i < len1 or j < len2:
            if j == len2 or i < len1 and ar1[i] < ar2[j]:
                ans[i + j] = ar1[i]
                i += 1
            else:
                ans[i + j] = ar2[j]
                j += 1
        return ans
