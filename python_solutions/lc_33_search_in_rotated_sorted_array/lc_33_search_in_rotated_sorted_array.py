class Solution:
    def search(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        for idx in range(len(nums)):
            if target == nums[idx]:
                return idx
        return -1

    def searchV1(self, nums: list[int], target: int) -> int:
        # Time complexity: O(2 log n). Space complexity: O(1).
        start, end = 0, len(nums) - 1
        boundary = self._find_boundary(nums)
        if nums[boundary] <= target <= nums[len(nums) - 1]:
            start, end = boundary, len(nums) - 1
        else:
            start, end = 0, boundary - 1
        return self._find_index(nums, target, start, end)

    def _find_index(self, nums: list[int], target: int, left: int, right: int):
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        return -1

    def _find_boundary(self, nums: list[int]):
        left, right = 0, len(nums) - 1
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] <= nums[-1]:
                right = mid
            else:
                left = mid + 1
        return left
