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

    def isLongPressedNameV2(self, name: str, typed: str) -> bool:
        # Time complexity: O(n + m). Space complexity: O(1).
        # Solution: https://leetcode.com/problems/long-pressed-name/editorial/
        # two pointers to the "name" and "typed" string respectively
        np, tp = 0, 0
        # advance two pointers, until we exhaust one of the strings
        while np < len(name) and tp < len(typed):
            if name[np] == typed[tp]:
                np += 1
                tp += 1
            elif tp >= 1 and typed[tp] == typed[tp - 1]:
                tp += 1
            else:
                return False
        # if there is still some characters left *unmatched* in the origin string,
        #   then we don't have a match.
        # e.g.  name = "abc"  typed = "aabb"
        if np != len(name):
            return False
        else:
            # In the case that there are some redundant characters left in typed
            # we could still have a match.
            # e.g.  name = "abc"  typed = "abccccc"
            while tp < len(typed):
                if typed[tp] != typed[tp - 1]:
                    return False
                tp += 1
        # both strings have been consumed
        return True
