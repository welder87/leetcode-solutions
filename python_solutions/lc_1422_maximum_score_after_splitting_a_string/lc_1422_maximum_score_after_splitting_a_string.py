class Solution:
    def maxScore(self, s: str) -> int:
        left = 1 if int(s[0]) == 0 else 0
        right = 0
        for i in range(1, len(s)):
            if int(s[i]) == 1:
                right += 1
        mx = left + right
        for i in range(1, len(s) - 1):
            if (val := int(s[i])) == 1:
                right -= 1
            elif val == 0:
                left += 1
            if (sm := left + right) > mx:
                mx = sm
        return mx
