import itertools


class Solution:
    def isSubstringPresent(self, s: str) -> bool:
        # Time complexity: O(n + n). Space complexity: O(m).
        # Method: Pair As String
        counter = set()
        for i in range(len(s) - 1, 0, -1):
            counter.add(f"{s[i]}{s[i - 1]}")
        for i in range(len(s) - 1):
            if f"{s[i]}{s[i + 1]}" in counter:
                return True
        return False

    def isSubstringPresentV1(self, s: str) -> bool:
        # Time complexity: O(n + n). Space complexity: O(m).
        # Method: Pair As Tuple
        counter = set()
        for i in range(len(s) - 1, 0, -1):
            counter.add((s[i], s[i - 1]))
        for i in range(len(s) - 1):
            if (s[i], s[i + 1]) in counter:
                return True
        return False

    def isSubstringPresentV2(self, s: str) -> bool:
        # Time complexity: O(n + n). Space complexity: O(m).
        # Method: Pair As Tuple + Itertools Builtins
        counter = set()
        for pair in itertools.pairwise(reversed(s)):
            counter.add(pair)
        for pair in itertools.pairwise(s):
            if pair in counter:
                return True
        return False

    def isSubstringPresentV3(self, s: str) -> bool:
        # Time complexity: O(n + n). Space complexity: O(m).
        # Method: Pair As Tuple + Itertools Builtins
        # Solution: https://leetcode.doocs.org/en/lc/3083/
        st = {(a, b) for a, b in itertools.pairwise(s[::-1])}
        return any((a, b) in st for a, b in itertools.pairwise(s))
