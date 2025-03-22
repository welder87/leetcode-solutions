class Solution:
    def searchInsert(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified itmo approach
        left, right = -1, len(nums)  # интервал
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                left = mid
            else:
                right = mid
        return left + 1

    def searchInsertV1(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified vanilla approach
        left, right = 0, len(nums) - 1  # отрезок
        while left <= right:
            mid = left + (right - left) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1
        return left

    def searchInsertV2(self, nums: list[int], target: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # algo.monster or yandex approach
        left, right = 0, len(nums)  # полуинтервал
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] >= target:
                right = mid
            else:
                left = mid + 1
        return left
