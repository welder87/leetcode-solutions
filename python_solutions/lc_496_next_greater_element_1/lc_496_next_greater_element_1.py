class Solution:
    def nextGreaterElement(self, nums1: list[int], nums2: list[int]) -> list[int]:
        # Time complexity: O(n * m). Space complexity: O(1).
        ans = [-1] * len(nums1)
        for idx, num in enumerate(nums1):
            cur, val = None, -1
            for nxt_num in nums2:
                if cur is None and num == nxt_num:
                    cur = nxt_num
                elif cur is not None and nxt_num > cur:
                    val = nxt_num
                    break
            ans[idx] = val
        return ans

    def nextGreaterElementV1(self, nums1: list[int], nums2: list[int]) -> list[int]:
        # Time complexity: O(n + m). Space complexity: O(k + p).
        # Solution based on: https://algo.monster/liteproblems/496
        stack = []
        h_table = {}
        for idx in range(len(nums2) - 1, -1, -1):
            while stack and stack[-1] < nums2[idx]:
                stack.pop()
            if stack:
                h_table[nums2[idx]] = stack[-1]
            stack.append(nums2[idx])
        return [h_table.get(num, -1) for num in nums1]
