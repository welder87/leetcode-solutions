from typing import Optional

from common_datastructures import ListNode


class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # привет
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow, fast = slow.next, fast.next.next
            if slow == fast:
                return slow.next
        return None
