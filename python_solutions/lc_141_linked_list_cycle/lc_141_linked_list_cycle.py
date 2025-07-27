from typing import Optional

from common_datastructures import ListNode


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        # Time complexity: O(n). Space complexity: O(1).
        slow, fast = head, head
        while slow is not None and fast is not None and fast.next is not None:
            slow, fast = slow.next, fast.next.next
            if slow == fast:
                return True
        return False
