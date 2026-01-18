from collections import Counter


class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        # Time complexity: O(n + m). Space complexity: O(k).
        counter = Counter(magazine)
        for char in ransomNote:
            try:
                cnt = counter[char]
            except KeyError:
                return False
            if cnt > 0:
                cnt -= 1
                counter[char] = cnt
            else:
                return False
        return True

    def canConstructV1(self, ransomNote: str, magazine: str) -> bool:
        # Time complexity: O(n + m). Space complexity: O(k).
        # Solution: https://leetcode.doocs.org/en/lc/383
        cnt = Counter(magazine)
        for c in ransomNote:
            cnt[c] -= 1
            if cnt[c] < 0:
                return False
        return True
