class Solution:
    def isSumEqual(self, firstWord: str, secondWord: str, targetWord: str) -> bool:
        # Time complexity: O(n + m + k). Space complexity: O(1).
        return calc_num(firstWord) + calc_num(secondWord) == calc_num(targetWord)


def calc_num(word: str) -> int:
    m = 1
    num = 0
    for char in reversed(word):
        num += (ord(char) - ord("a")) * m
        m *= 10
    return num
