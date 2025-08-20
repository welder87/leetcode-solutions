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

    def sortArrayV1(self, nums: list[int]) -> list[int]:
        # Time complexity: O(n log n). Space complexity: O(2n).
        # Method: Recursive Merge Sort With Single Buffer.
        buffer = [0] * len(nums)
        self._merge_sort_v1(nums, buffer, 0, len(nums) - 1)
        return nums

    def _merge_sort_v1(self, nums: list[int], buffer: list[int], left: int, right: int):
        len = right - left + 1
        if len < 2:
            return nums
        mid = (right + left) // 2
        self._merge_sort_v1(nums, buffer, left, mid)
        self._merge_sort_v1(nums, buffer, mid + 1, right)
        self._merge_two_arrays_v1(nums, buffer, left, mid, mid + 1, right)

    def _merge_two_arrays_v1(
        self,
        nums: list[int],
        buffer: list[int],
        left1: int,
        right1: int,
        left2: int,
        right2: int,
    ) -> None:
        i, j, n = left1, left2, left1
        while i <= right1 or j <= right2:
            if j > right2 or i <= right1 and nums[i] < nums[j]:
                buffer[n] = nums[i]
                i += 1
            else:
                buffer[n] = nums[j]
                j += 1
            n += 1
        for n in range(left1, right2 + 1):
            nums[n] = buffer[n]
