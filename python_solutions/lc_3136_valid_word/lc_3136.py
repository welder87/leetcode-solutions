class Solution:
    def isValid(self, word: str) -> bool:
        # Time complexity O(n), Space complexity O(1)
        if len(word) < 3:
            return False
        has_vowel = False
        has_consonant = False
        for char in word:
            is_digit = char.isdigit()
            if not is_digit and not char.isupper() and not char.islower():
                return False
            if is_digit:
                continue
            char = char.lower()
            if char == "a" or char == "e" or char == "i" or char == "o" or char == "u":
                has_vowel = True
            else:
                has_consonant = True
        return has_vowel and has_consonant
