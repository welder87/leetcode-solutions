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

    def deleteDuplicatesV2(self, head: ListNode) -> ListNode:
        # Time complexity: O(n). Space complexity: O(1).
        # Solution: https://algo.monster/liteproblems/83
        # Initialize current to point to the head of the list
        current = head
        # Traverse the linked list
        while current and current.next:
            # If the current value is equal to the value in the next node
            if current.val == current.next.val:
                # Bypass the next node as it's a duplicate
                current.next = current.next.next
            else:
                # Move to the next unique value if no duplicate is found
                current = current.next
        # Return the head of the updated list
        return head
