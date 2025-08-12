from typing import Optional

from common_datastructures import ListNode


class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        # Method: Data Overwriting.
        if head is None:
            return head
        left, right = head, head.next
        while right is not None:
            left.val, right.val = right.val, left.val
            left = right.next
            if left is None:
                break
            right = left.next
        return head

    def swapPairsV1(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        # Method: Swapping node using pointers.
        dummy = ListNode(val=-1, next=head)
        prev, head = dummy, head
        while head is not None and head.next is not None:
            nxt = head.next
            prev.next = nxt
            head.next = nxt.next
            nxt.next = head
            prev = head
            head = head.next
        return dummy.next
