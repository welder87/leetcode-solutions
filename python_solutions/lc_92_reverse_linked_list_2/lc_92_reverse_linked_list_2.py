from typing import Optional

from common_datastructures import ListNode


class Solution:
    def reverseBetween(
        self,
        head: Optional[ListNode],
        left: int,
        right: int,
    ) -> Optional[ListNode]:
        # Time complexity: O(n+m). Space complexity: O(1).
        cur = head
        counter = 0
        pre_start, start, end, post_end, prev = None, None, None, None, None
        while cur is not None:
            counter += 1
            if counter == left:
                start = cur
                pre_start = prev
            if counter == right:
                end = cur
                post_end = cur.next
                break
            prev = cur
            cur = cur.next

        if pre_start is not None:
            pre_start.next = None

        if end is not None:
            end.next = None

        prev, cur = None, start
        while cur is not None:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt

        if pre_start is None and post_end is None:
            return end
        if pre_start is None and post_end is not None:
            start.next = post_end
            return end
        if pre_start is not None and post_end is None:
            pre_start.next = end
            return head

        pre_start.next = end
        start.next = post_end

        return head
