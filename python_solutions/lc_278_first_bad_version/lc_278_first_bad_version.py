# The isBadVersion API is already defined for you.


def isBadVersion(version: int) -> bool:
    pass


class Solution:
    def firstBadVersion(self, n: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified yandex / algo.monster / leetcode
        left, right = 1, n
        while left < right:
            mid = left + (right - left) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid + 1
        return left

    def firstBadVersionV1(self, n: int) -> int:
        # Time complexity: O(log n). Space complexity: O(1).
        # modified itmo
        left, right = 0, n + 1
        while right > left + 1:
            mid = left + (right - left) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid
        return left + 1
