from typing import Optional

from common_datastructures import ListNode


class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n/2 + 1). Space complexity: O(1).
        slow, fast, prev = head, head, None
        while fast is not None and fast.next is not None:
            prev = slow
            slow, fast = slow.next, fast.next.next
        if prev is None:
            return None
        prev.next = slow.next
        slow.next = None
        return head
