from bisect import bisect_right, bisect_left


class Solution:
    def search(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # vanilla
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1
        return -1

    def searchV1(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified itmo
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                left = mid
            else:
                right = mid
        return -1

    def searchV2(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified yandex
        left, right = 0, len(nums)
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] > target:
                right = mid
            else:
                left = mid + 1
        if left > 0 and nums[left - 1] == target:
            return left - 1
        return -1

    def searchV3(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # built-in bisect_right
        ans = bisect_right(nums, target)
        if ans > 0 and nums[ans - 1] == target:
            return ans - 1
        return -1

    def searchV4(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # built-in bisect_left
        ans = bisect_left(nums, target)
        if ans < len(nums) and nums[ans] == target:
            return ans
        return -1
