from typing import Optional

from common_datastructures import ListNode


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(n).
        if head is None:
            return None
        new_head = ListNode(val=head.val)
        prev = new_head
        cur = head.next
        while cur is not None:
            if cur.val != prev.val:
                prev.next = ListNode(val=cur.val)
                prev = prev.next
            cur = cur.next
        return new_head

    def deleteDuplicatesV1(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        if head is None:
            return None
        prev = head
        cur = head.next
        while cur is not None:
            if cur.val == prev.val:
                prev.next = cur.next
                cur.next = None
                cur = prev.next
            else:
                prev = cur
                cur = cur.next
        return head
