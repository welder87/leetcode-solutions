class Solution:
    def findPeakElement(self, nums: list[int]) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        if len(nums) == 1:
            return 0
        if len(nums) == 2:
            return 0 if nums[0] > nums[1] else 1

        for i in range(-1, len(nums)):
            if (
                (i == -1 and nums[i + 1] > nums[i + 2])
                or (i + 2 == len(nums) and nums[i] < nums[i + 1])
                or (nums[i] < nums[i + 1] > nums[i + 2])
            ):
                return i + 1
        return -1

    def findPeakElementV1(self, nums: list[int]) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        if len(nums) == 1:
            return 0

        left, right = -1, len(nums)
        while 1:
            mid = left + (right - left) // 2
            if (
                (mid - 1 == -1 and nums[mid] > nums[mid + 1])
                or (mid + 1 == len(nums) and nums[mid - 1] < nums[mid])
                or (nums[mid - 1] < nums[mid] > nums[mid + 1])
            ):
                return mid
            if nums[mid] > nums[mid + 1]:
                right = mid
            else:
                left = mid
        return -1

    def findPeakElementV2(self, nums: list[int]) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        left, right = 0, len(nums) - 1
        while right > left:
            mid = left + (right - left) // 2
            if nums[mid] > nums[mid + 1]:
                right = mid
            else:
                left = mid + 1
        return right
