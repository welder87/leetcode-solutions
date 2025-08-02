from typing import Optional

from common_datastructures import ListNode


class Solution:
    def getIntersectionNode(
        self,
        headA: ListNode,
        headB: ListNode,
    ) -> Optional[ListNode]:
        # Time complexity: O(m + n). Space complexity: O(m).
        unique_ids = set()
        cur = headA
        while cur is not None:
            unique_ids.add(id(cur))
            cur = cur.next
        cur = headB
        while cur is not None and id(cur) not in unique_ids:
            cur = cur.next
        return cur

    def getIntersectionNodeV1(
        self,
        headA: ListNode,
        headB: ListNode,
    ) -> Optional[ListNode]:
        # Time complexity: O(2*m + 2*n). Space complexity: O(1).
        h1, h2 = headA, headB
        l1, l2 = 0, 0
        while h1 is not None:
            l1 += 1
            h1 = h1.next
        while h2 is not None:
            l2 += 1
            h2 = h2.next

        if l1 >= l2:
            h1, h2 = headA, headB
        else:
            h1, h2 = headB, headA

        p = abs(l1 - l2)
        cnt = 0

        while h1 is not None and p != cnt:
            cnt += 1
            h1 = h1.next

        while h1 is not None and h2 is not None and id(h1) != id(h2):
            h1 = h1.next
            h2 = h2.next
        return h1

    def getIntersectionNodeV2(
        self,
        headA: ListNode,
        headB: ListNode,
    ) -> Optional[ListNode]:
        # Time complexity: O(m + n). Space complexity: O(1).
        l1, l2 = headA, headB
        while l1 != l2:
            l1 = l1.next if l1 else headB
            l2 = l2.next if l2 else headA
        return l1
