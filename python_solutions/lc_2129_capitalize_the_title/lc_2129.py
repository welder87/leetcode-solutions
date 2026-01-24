class Solution:
    def capitalizeTitle(self, title: str) -> str:
        # Time complexity O(n + n + n)
        # Space complexity O(n + n)
        ans = list(title)
        start = 0
        for i in range(len(ans)):
            if ans[i] != " ":
                ans[i] = ans[i].lower()
                continue
            if i - start > 2:
                ans[start] = ans[start].upper()
            start = i + 1
        if len(ans) - start > 2:
            ans[start] = ans[start].upper()
        return "".join(ans)
