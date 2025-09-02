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
