from collections import Counter


class Solution:
    def firstUniqChar(self, s: str) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        counter = Counter(s)
        for idx, char in enumerate(s):
            if counter[char] == 1:
                return idx
        return -1

    def firstUniqCharV1(self, s: str) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        counter = [0] * 26
        diff = ord("a")
        for char in s:
            counter[ord(char) - diff] += 1
        for idx, char in enumerate(s):
            if counter[ord(char) - diff] == 1:
                return idx
        return -1
