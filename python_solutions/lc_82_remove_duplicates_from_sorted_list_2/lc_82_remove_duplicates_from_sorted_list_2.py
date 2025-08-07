from typing import Optional

from common_datastructures import ListNode


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None

        right = head.next
        while right is not None and head.val == right.val:
            head.next = None
            head, right = right, right.next
        if right is None:
            head = None
            return None
        head.next = None
        head = right

        # -1, -1, -1, 0, 2, 3, 3, 3, 4, 4, 5, 5, 5
        prev, cur, nxt = head, head, head.next

        while nxt is not None:
            while nxt is not None and cur.val == nxt.val:
                if prev is not None:
                    prev.next = nxt
                cur.next = None
                cur, nxt = nxt, nxt.next

            if nxt is not None:
                cur, nxt = nxt, nxt.next

            prev, cur, nxt = cur, nxt, nxt.next

        return head
