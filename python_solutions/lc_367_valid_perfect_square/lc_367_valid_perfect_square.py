class Solution(object):
    def isPerfectSquare(self, num: int) -> bool:
        """
        :type num: int
        :rtype: bool
        """
        left, right = 1, (num // 2)
        while left + 1 < right:
            mid = left + (right - left) // 2
            if mid * mid < num:
                left = mid
            else:
                right = mid
        if left * left == num or right * right == num:
            return True
        return False
