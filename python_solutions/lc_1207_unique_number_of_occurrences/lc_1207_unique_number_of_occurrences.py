from collections import Counter


class Solution:
    def uniqueOccurrences(self, arr: list[int]) -> bool:
        counter = Counter(arr)
        return len(counter.values()) == len(set(counter.values()))

    def uniqueOccurrencesV1(self, arr: list[int]) -> bool:
        counter = Counter(arr)
        unique_cnt = set()
        for _, cnt in counter.items():
            if cnt in unique_cnt:
                return False
            unique_cnt.add(cnt)
        return True
