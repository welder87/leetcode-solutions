from collections import Counter


class Solution:
    def findLucky(self, arr: list[int]) -> int:
        # Time complexity: O(n + m). Space complexity: O(m).
        largest_lucky = -1
        for num, cnt in Counter(arr).items():
            if num == cnt and largest_lucky < num:
                largest_lucky = num
        return largest_lucky

    def findLuckyV1(self, arr: list[int]) -> int:
        # Time complexity: O(2n + m). Space complexity: O(m).
        counter = [0] * (max(arr) + 1)
        for num in arr:
            counter[num] += 1
        largest_lucky = -1
        for num in range(1, len(counter)):
            if num == counter[num] and largest_lucky < num:
                largest_lucky = num
        return largest_lucky
