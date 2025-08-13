from typing import Optional

from common_datastructures import ListNode


class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        dummy_lt, dummy_ge = ListNode(val=-1000), ListNode(val=1000)
        lt, ge = dummy_lt, dummy_ge
        while head is not None:
            spam = head.next
            head.next = None
            if head.val >= x:
                ge.next = head
                ge = ge.next
            else:
                lt.next = head
                lt = lt.next
            head = spam
        lt.next = dummy_ge.next
        return dummy_lt.next
