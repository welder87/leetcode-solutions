class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        # Time complexity O(sum(char)) + O(sum(char))
        # Space complexity O(m + k + p)
        words = s.split(" ")
        if len(words) != len(pattern):
            return False
        counter1 = {}
        counter2 = {}
        for word, char in zip(words, pattern):
            existing_word = counter1.get(char)
            existing_char = counter2.get(word)
            if existing_word is None and existing_char is None:
                counter1[char] = word
                counter2[word] = char
                continue
            if existing_word != word or existing_char != char:
                return False
        return True
