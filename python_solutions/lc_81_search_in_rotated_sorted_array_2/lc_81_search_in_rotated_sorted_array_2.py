class Solution:
    def search(self, nums: list[int], target: int) -> bool:
        # Time complexity: O(log n). Space complexity: O(1).
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] > nums[right]:
                # тогда отрезок [left, mid] отсортирован
                if nums[left] <= target <= nums[mid]:
                    # тогда target на отрезке [left, mid]
                    right = mid
                else:
                    # тогда target на отрезке [mid + 1, right]
                    left = mid + 1
            elif nums[mid] < nums[right]:
                # тогда отрезок [mid + 1, right] отсортирован
                if nums[mid] < target <= nums[right]:
                    # тогда target на отрезке [mid + 1, right]
                    left = mid + 1
                else:
                    # тогда target на отрезке [left, mid]
                    right = mid
            else:
                # тогда нельзя точно сказать на каком отрезке лежит искомое значение
                right -= 1
        return nums[left] == target
