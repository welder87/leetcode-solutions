class Solution:
    def checkIfExist(self, arr: list[int]) -> bool:
        # Time complexity: O(2 * n log n). Space complexity: O(n).
        # modified itmo
        arr.sort()
        for target_index, target in enumerate(arr):
            left, right = -1, len(arr)
            while right > left + 1:
                mid = left + (right - left) // 2
                val = 2 * arr[mid]
                if target_index != mid and val == target:
                    return True
                if val < target:
                    left = mid
                else:
                    right = mid
        return False
