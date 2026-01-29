from itertools import zip_longest


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

    def longestCommonPrefixV1(self, strs: list[str]) -> str:
        # Time complexity: O(sum(char)). Space complexity: O(1).
        if len(strs) == 1:
            return strs[0]
        ans = list(strs[0])
        for i in range(1, len(strs)):
            for j, (first, second) in enumerate(zip_longest(strs[i], ans)):
                if first != second:
                    ans = ans[:j]
                    break
        return "".join(ans)

    def longestCommonPrefixV2(self, strs: list[str]) -> str:
        # Time complexity: O(sum(char)). Space complexity: O(1).
        ans = strs[0]
        for i in range(1, len(strs)):
            for j, (first, second) in enumerate(zip_longest(strs[i], ans)):
                if first != second:
                    ans = ans[:j]
                    break
        return ans
