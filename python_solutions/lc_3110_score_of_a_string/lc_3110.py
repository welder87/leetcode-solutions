class Solution:
    def scoreOfString(self, s: str) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        sum_ = 0
        for i in range(len(s) - 1):
            sum_ += abs(ord(s[i]) - ord(s[i + 1]))
        return sum_
