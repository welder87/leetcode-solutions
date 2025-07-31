from typing import Optional

from common_datastructures import ListNode


class Solution:
    def getIntersectionNode(
        self,
        headA: ListNode,
        headB: ListNode,
    ) -> Optional[ListNode]:
        # Time complexity: O(2*n). Space complexity: O(m).
        unique_ids = set()
        cur = headA
        while cur is not None:
            unique_ids.add(id(cur))
            cur = cur.next
        cur = headB
        while cur is not None and id(cur) not in unique_ids:
            cur = cur.next
        return cur
