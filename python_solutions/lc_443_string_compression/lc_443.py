class Solution:
    def compress(self, chars: list[str]) -> int:
        chars.sort()
        i, j, n = 0, 1, 0
        while i < len(chars):
            counter = 1
            while j < len(chars) and chars[i] == chars[j]:
                j += 1
                counter += 1
            while n < len(chars):
                chars
            i = j
            j += 1
