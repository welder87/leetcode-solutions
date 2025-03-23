class Solution:
    def intersection(self, nums1: list[int], nums2: list[int]) -> list[int]:
        # Time complexity: O(n+m+k). Space complexity: O(n+m).
        result = set(nums1)
        result.intersection_update(nums2)
        return list(result)

    def intersection_v1(self, nums1: list[int], nums2: list[int]) -> list[int]:
        # Time complexity: O(n log n + m log m + m + n). Space complexity: O(n+m).
        nums1.sort()
        nums2.sort()
        i, j, l1, l2, result = 0, 0, len(nums1), len(nums2), set()
        while i < l1 and j < l2:
            if nums1[i] == nums2[j]:
                result.add(nums1[i])
                i += 1
                j += 1
            elif nums1[i] < nums2[j]:
                i += 1
            else:
                j += 1
        return list(result)

    def intersection_v2(self, nums1: list[int], nums2: list[int]) -> list[int]:
        # Time complexity: O(n log n + m log m). Space complexity: O(n + m + k).
        # itmo binary search
        result = set()
        nums1.sort()
        for target in nums2:
            left, right = -1, len(nums1)
            while right > left + 1:
                mid = left + (right - left) // 2
                if nums1[mid] == target:
                    result.add(target)
                    break
                if nums1[mid] > target:
                    right = mid
                else:
                    left = mid
        return list(result)
