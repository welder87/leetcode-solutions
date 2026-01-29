class Solution:
    def longestCommonPrefix(self, strs: list[str]) -> str:
        # Time complexity: O(sum(char)). Space complexity: O(k).
        if not len(strs):
            return ""
        if len(strs) == 1:
            return strs[0]
        prefix = ""
        for symbols in zip(*strs):
            r = set(symbols)
            if len(r) != 1:
                break
            prefix = f"{prefix}{r.pop()}"
        return prefix
