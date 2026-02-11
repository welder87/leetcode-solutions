class Solution:
    def isSumEqual(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        # Time complexity: O(n + m + k). Space complexity: O(1).
        return calc_num(firstWord) + calc_num(secondWord) == calc_num(targetWord)

    def isSumEqualV1(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        # Time complexity: O(n + m + k). Space complexity: O(1).
        # Solution: https://leetcode.doocs.org/en/lc/1880 .
        def f(s: str) -> int:
            ans, a = 0, ord("a")
            for c in map(ord, s):
                x = c - a
                ans = ans * 10 + x
            return ans

        return f(firstWord) + f(secondWord) == f(targetWord)


def calc_num(word: str) -> int:
    m = 1
    num = 0
    for char in reversed(word):
        num += (ord(char) - ord("a")) * m
        m *= 10
    return num
