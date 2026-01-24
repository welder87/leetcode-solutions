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

    def capitalizeTitleV1(self, title: str) -> str:
        # Time complexity O(n + n + n)
        # Space complexity O(n + n + n)
        return " ".join(
            word.title() if len(word) > 2 else word.lower() for word in title.split(" ")
        )

    def capitalizeTitleV2(self, title: str) -> str:
        # Time complexity O(n + n + n)
        # Space complexity O(n + n + n)
        # Solution: https://leetcode.doocs.org/en/lc/2129
        words = [w.lower() if len(w) < 3 else w.capitalize() for w in title.split()]
        return " ".join(words)
