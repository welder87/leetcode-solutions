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

    def shortestToCharV1(self, s: str, c: str) -> list[int]:
        # Time complexity: O(2*n). Space complexity: O(n).
        #  Solution: https://algo.monster/liteproblems/821
        n = len(s)
        # Initialize result array with maximum possible distance
        result = [n] * n
        # First pass: left to right
        # Track the most recent position of character c seen so far
        prev_c_position = float("-inf")
        for i in range(n):
            if s[i] == c:
                # Update the position of the most recent c
                prev_c_position = i
            # Calculate distance from current position to previous c
            result[i] = min(result[i], i - prev_c_position)
        # Second pass: right to left
        # Track the most recent position of character c seen from the right
        next_c_position = float("inf")
        for i in range(n - 1, -1, -1):
            if s[i] == c:
                # Update the position of the most recent c from right
                next_c_position = i
            # Calculate distance from current position to next c
            # and keep the minimum distance
            result[i] = min(result[i], next_c_position - i)
        return result
