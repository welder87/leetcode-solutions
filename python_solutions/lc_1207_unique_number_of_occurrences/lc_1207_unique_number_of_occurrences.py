from collections import Counter


class Solution:
    def uniqueOccurrences(self, arr: list[int]) -> bool:
        counter = Counter(arr)
        return len(counter.values()) == len(set(counter.values()))
