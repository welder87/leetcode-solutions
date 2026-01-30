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

    def longestCommonPrefixV3(self, strs: list[str]) -> str:
        # Time complexity: O(sum(char)). Space complexity: O(1).
        # Method: Horizontal scanning
        # Solution: https://leetcode.com/problems/longest-common-prefix/editorial/#approach-1-horizontal-scanning
        if len(strs) == 0:
            return ""
        prefix = strs[0]
        for i in range(1, len(strs)):
            while strs[i].find(prefix) != 0:
                prefix = prefix[0 : len(prefix) - 1]
                if prefix == "":
                    return ""
        return prefix

    def longestCommonPrefixV4(self, strs: list[str]) -> str:
        # Time complexity: O(sum(char)). Space complexity: O(1).
        # Method: Vertical scanning
        # Solution: https://leetcode.com/problems/longest-common-prefix/editorial/#approach-2-vertical-scanning
        if len(strs) == 0:
            return ""
        for i in range(len(strs[0])):
            c = strs[0][i]
            for j in range(1, len(strs)):
                if i == len(strs[j]) or strs[j][i] != c:
                    return strs[0][0:i]
        return strs[0]
