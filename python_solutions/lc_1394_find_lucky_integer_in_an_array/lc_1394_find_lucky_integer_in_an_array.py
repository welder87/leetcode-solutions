from collections import Counter


class Solution:
    def findLucky(self, arr: list[int]) -> int:
        # Time complexity: O(n + m). Space complexity: O(m).
        largest_lucky = -1
        for num, cnt in Counter(arr).items():
            if num == cnt and largest_lucky < num:
                largest_lucky = num
        return largest_lucky
