class Solution:
    def shortestToChar(self, s: str, c: str) -> list[int]:
        # Time complexity: O(2*n). Space complexity: O(n).
        ans = [0] * len(s)
        val, left, right = 0, -1, -1
        # left to right
        for i in range(len(s)):
            if s[i] == c:
                val = 1
                if left == -1:
                    left = i
                right = i
            else:
                ans[i] = val
                val += 1
        val = 0
        # right to left
        for i in range(right, -1, -1):
            if s[i] == c:
                val = 1
            else:
                ans[i] = min(ans[i], val)
                val += 1
        # left check
        for i in range(left):
            ans[i] = left - i
        return ans
