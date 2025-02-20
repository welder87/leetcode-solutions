class Solution:
    def findShortestSubArray(self, nums: list[int]) -> int:
        # Time complexity: O(n + m). Space complexity: O(m).
        map_items = {}
        for idx, num in enumerate(nums):
            if num not in map_items:
                map_items[num] = [1, [idx, None]]
            else:
                map_items[num][0] += 1
                map_items[num][1][1] = idx
        max_cnt, degree = None, None
        for cnt, (first_index, last_index) in map_items.values():
            cur = 1 if last_index is None else last_index - first_index + 1
            if max_cnt is None:
                max_cnt = cnt
                degree = cur
            elif max_cnt == cnt:
                degree = cur if degree > cur else degree
            elif max_cnt < cnt:
                max_cnt = cnt
                degree = cur
        return degree
