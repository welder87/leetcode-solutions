from typing import Optional

from common_datastructures import ListNode


class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n/2 + 1). Space complexity: O(1).
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow, fast = slow.next, fast.next.next
        return slow
