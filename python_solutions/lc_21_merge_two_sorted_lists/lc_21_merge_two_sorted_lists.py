from typing import Optional

from common_datastructures import ListNode


class Solution:
    def mergeTwoLists(
        self,
        list1: Optional[ListNode],
        list2: Optional[ListNode],
    ) -> Optional[ListNode]:
        # Time complexity: O(n+m). Space complexity: O(1).
        if list1 is None and list2 is None:
            return None
        if list1 is None:
            return list2
        if list2 is None:
            return list1

        if list1.val < list2.val:
            head = ListNode(val=list1.val)
            p1, p2 = list1.next, list2
        else:
            head = ListNode(val=list2.val)
            p1, p2 = list1, list2.next
        cur = head

        while p1 is not None or p2 is not None:
            if p2 is None or p1 is not None and p1.val < p2.val:
                cur.next = ListNode(val=p1.val)
                p1 = p1.next
            else:
                cur.next = ListNode(val=p2.val)
                p2 = p2.next
            cur = cur.next
        return head

    def mergeTwoListsV1(
        self,
        list1: Optional[ListNode],
        list2: Optional[ListNode],
    ) -> Optional[ListNode]:
        # Time complexity: O(n+m). Space complexity: O(1).
        dummy = node = ListNode(val=0)
        while list1 is not None or list2 is not None:
            if list2 is None or list1 is not None and list1.val < list2.val:
                node.next = list1
                list1 = list1.next
            else:
                node.next = list2
                list2 = list2.next
            node = node.next
            node.next = None
        return dummy.next
