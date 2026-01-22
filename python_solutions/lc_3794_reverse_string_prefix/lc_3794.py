class Solution:
    def reversePrefix(self, s: str, k: int) -> str:
        # Time complexity O(n + n), space complexity O(n + n)
        # Method: Two-pointer approach (converging pointers).
        ans = list(s)
        i, j = 0, k - 1
        while i < j:
            ans[i], ans[j] = ans[j], ans[i]
            i += 1
            j -= 1
        return "".join(ans)
