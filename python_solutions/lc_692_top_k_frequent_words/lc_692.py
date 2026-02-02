from collections import Counter


class Solution:
    def topKFrequent(self, words: list[str], k: int) -> list[str]:
        # Time complexity O(sum(char)) + O(n log sum(char)) + O(m)
        # Space complexity O(k) + O(z) + O(m)
        # Method: Hash Table + Sorting
        counter = Counter(words)
        inter = sorted(
            (item for item in counter.items()),
            key=lambda x: (-x[1], x[0]),
        )
        ans = [""] * k
        for i, (key, _) in enumerate(inter):
            if i >= k:
                break
            ans[i] = key
        return ans

    def topKFrequentV1(self, words: list[str], k: int) -> list[str]:
        # Time complexity O(sum(char)) + O(n log sum(char)) + O(m)
        # Space complexity O(k) + O(z) + O(m)
        # Method: Hash Table + Sorting
        # Solution: https://leetcode.doocs.org/en/lc/692
        cnt = Counter(words)
        return sorted(cnt, key=lambda x: (-cnt[x], x))[:k]
