class Solution:
    def findKthPositive(self, arr: list[int], k: int) -> int:
        # Time complexity: O(n log n). Space complexity: O(1).
        # O(arr[−1]∗K∗log(len(arr)))
        min_num, max_num = 1, arr[-1] * k + 1
        counter = 0
        num = -1
        for num in range(min_num, max_num + 1, 1):
            if not self.bs(arr, num):
                counter += 1
            if counter == k:
                break
        return num

    @staticmethod
    def bs(nums: list[int], target: int) -> bool:
        left, right = 0, len(nums) - 1
        while left + 1 < right:
            mid = left + (right - left) // 2
            if nums[mid] < target:
                left = mid
            else:
                right = mid
        if nums[left] == target or nums[right] == target:
            return True
        return False

    def findKthPositiveV1(self, arr: list[int], k: int) -> int:
        # Time complexity: O(2 n). Space complexity: O(n).
        s = set(arr)
        min_num, max_num = 1, arr[-1] * k + 1
        counter = 0
        num = -1
        for num in range(min_num, max_num + 1, 1):
            if num not in s:
                counter += 1
            if counter == k:
                break
        return num
