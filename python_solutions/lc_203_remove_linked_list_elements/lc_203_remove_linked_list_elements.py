from typing import Optional

from common_datastructures import ListNode


class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        while head is not None and head.val == val:
            head = head.next
        if head is None:
            return None
        cur = head.next
        prev = head
        while cur is not None:
            if cur.val == val:
                prev.next = cur.next
                cur.next = None
                cur = prev.next
            else:
                prev = cur
                cur = cur.next
        return head
