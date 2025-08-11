from typing import Optional

from common_datastructures import ListNode


class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        if head is None:
            return head
        left, right = head, head.next
        while right is not None:
            left.val, right.val = right.val, left.val
            left = right.next
            if left is None:
                break
            right = left.next
        return head
