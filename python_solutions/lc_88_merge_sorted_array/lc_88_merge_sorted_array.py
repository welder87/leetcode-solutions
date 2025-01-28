class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        i, j = m - 1, n - 1
        while i >= 0 or j >= 0:
            if i < 0 or (j >= 0 and nums1[i] <= nums2[j]):
                nums1[i + j + 1] = nums2[j]
                j -= 1
            else:
                nums1[i + j + 1] = nums1[i]
                i -= 1

    def merge_v1(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        # Time complexity: O(n log n). Space complexity: O(n).
        for i in range(n):
            nums1[m + i] = nums2[i]
        nums1.sort()
