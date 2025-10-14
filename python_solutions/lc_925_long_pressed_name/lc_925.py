import itertools


class Solution:
    def isLongPressedName(self, name: str, typed: str) -> bool:
        # Time complexity: O(n + m). Space complexity: O(1).
        i, j = 0, 0
        while i < len(name) and j < len(typed) and typed[j] == name[i]:
            counter1, counter2 = 0, 0
            sym = name[i]
            while i < len(name) and sym == name[i]:
                i += 1
                counter1 += 1
            while j < len(typed) and typed[j] == sym:
                j += 1
                counter2 += 1
            if counter1 > counter2:
                return False
        return i == len(name) and j == len(typed)

    def isLongPressedNameV1(self, name, typed):
        # Time complexity: O(n + m). Space complexity: O(n + m).
        # Solution: https://leetcode.com/problems/long-pressed-name/editorial/
        g1 = [(k, len(list(grp))) for k, grp in itertools.groupby(name)]
        g2 = [(k, len(list(grp))) for k, grp in itertools.groupby(typed)]
        if len(g1) != len(g2):
            return False

        return all(k1 == k2 and v1 <= v2 for (k1, v1), (k2, v2) in zip(g1, g2))
