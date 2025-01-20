from collections import defaultdict


class Solution:
    def mostCommonWord(self, paragraph: str, banned: list[str]) -> str:
        paragraph = paragraph.lower()
        start, end = ord("a"), ord("z")
        unique_words = defaultdict(int)
        banned = set(banned)
        word_symbols = []
        for symbol in paragraph:
            if start <= ord(symbol) <= end:
                word_symbols.append(symbol)
            elif word_symbols:
                word = "".join(word_symbols)
                if word not in banned:
                    unique_words[word] += 1
                word_symbols.clear()
        if word_symbols:
            word = "".join(word_symbols)
            if word not in banned:
                unique_words[word] += 1
            word_symbols.clear()
        if not unique_words:
            return ""
        return max(unique_words.items(), key=lambda x: x[1])[0]
