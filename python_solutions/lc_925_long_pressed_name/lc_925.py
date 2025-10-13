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
