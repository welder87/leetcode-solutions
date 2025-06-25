from typing import Literal


class Solution(object):
    def guessNumber(self, n: int) -> int:
        """
        :type n: int
        :rtype: int
        """
        left, right = 1, n
        while left + 1 < right:
            mid = left + (right - left) // 2
            if guess(mid) > 0:
                left = mid
            else:
                right = mid
        if guess(left) == 0:
            return left
        return right

    def guessNumberV1(self, n: int) -> int:
        """
        :type n: int
        :rtype: int
        """
        left, right = 1, n
        while left <= right:
            mid = left + (right - left) // 2
            if (res := guess(mid)) == 0:
                return mid
            elif res > 0:
                left = mid + 1
            else:
                right = mid - 1
        return left


def guess(num: int) -> Literal[0, -1, 1]: ...
