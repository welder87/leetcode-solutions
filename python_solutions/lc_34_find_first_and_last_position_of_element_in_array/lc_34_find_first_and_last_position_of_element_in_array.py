import operator


class Solution:
    def searchRange(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2 log n). Space complexity: O(1).
        # modified itmo
        if len(nums) == 0:
            return [-1, -1]
        left, right = self.left_bs(nums, target), self.right_bs(nums, target)
        return [-1, -1] if left > right else [left, right]

    def right_bs(self, nums: list[int], target: int) -> int:
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] <= target:
                left = mid
            else:
                right = mid
        return left

    def left_bs(self, nums: list[int], target: int) -> int:
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid
            else:
                right = mid
        return right

    def searchRangeV1(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2 log n). Space complexity: O(1).
        # modified itmo
        if len(nums) == 0:
            return [-1, -1]
        left, right = self.bs(nums, target, "left"), self.bs(nums, target, "right")
        return [-1, -1] if left > right else [left, right]

    def bs(self, nums: list[int], target: int, kind: str) -> int:
        left, right = -1, len(nums)
        compare = operator.lt if kind == "left" else operator.le
        while right > left + 1:
            mid = left + (right - left) // 2
            if compare(nums[mid], target):
                left = mid
            else:
                right = mid
        return right if kind == "left" else left

    def searchRangeV2(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2 log n). Space complexity: O(1).
        # modified itmo
        if len(nums) == 0:
            return [-1, -1]
        left = self._bs(nums, target, "left", -1, len(nums))
        if left == -1:
            return [-1, -1]
        right = self._bs(nums, target, "right", left - 1, len(nums))
        return [-1, -1] if left > right else [left, right]

    def _bs(
        self,
        nums: list[int],
        target: int,
        kind: str,
        left: int,
        right: int,
    ) -> int:
        compare = operator.lt if kind == "left" else operator.le
        while right > left + 1:
            mid = left + (right - left) // 2
            if compare(nums[mid], target):
                left = mid
            else:
                right = mid
        return right if kind == "left" else left

    def searchRangeV3(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2 log n). Space complexity: O(1).
        # modified yandex
        if len(nums) == 0:
            return [-1, -1]
        left = self.yandex_left(nums, target)
        right = self.yandex_right(nums, target)
        return [-1, -1] if left > right else [left, right]

    def yandex_right(self, nums: list[int], target: int) -> int:
        left, right = 0, len(nums)
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] <= target:
                left = mid + 1
            else:
                right = mid
        return left - 1

    def yandex_left(self, nums: list[int], target: int) -> int:
        left, right = 0, len(nums)
        while left < right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid
        return left
