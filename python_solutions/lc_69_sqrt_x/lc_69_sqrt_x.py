class Solution(object):
    def mySqrt(self, x: int) -> int:
        """
        :type x: int
        :rtype: int
        """
        left, right = 0, x
        while left + 1 < right:
            mid = left + (right - left) // 2
            if mid * mid <= x:
                left = mid
            else:
                right = mid
        if right * right == x:
            return right
        return left
