from typing import Optional

from common_datastructures import ListNode


class Solution:
    def mergeTwoLists(
        self,
        list1: Optional[ListNode],
        list2: Optional[ListNode],
    ) -> Optional[ListNode]:
        # Time complexity: O(n+m). Space complexity: O(n+m).
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

    def mergeTwoListsV2(self, list1: ListNode, list2: ListNode) -> ListNode:
        # Time complexity: O(n+m). Space complexity: O(1).
        # Solution: https://algo.monster/liteproblems/21
        # Creating a sentinel node which helps to easily return the head of the merged list
        sentinel = ListNode(val=0)
        # Current node is used to keep track of the end of the merged list
        current = sentinel
        # Iterate while both lists have nodes
        while list1 and list2:
            # Choose the smaller value from either list1 or list2
            if list1.val <= list2.val:
                current.next = list1  # Append list1 node to merged list
                list1 = list1.next  # Move to the next node in list1
            else:
                current.next = list2  # Append list2 node to merged list
                list2 = list2.next  # Move to the next node in list2
            # Move the current pointer forward in the merged list
            current = current.next
        # Add any remaining nodes from list1 or list2 to the merged list
        # If one list is fully traversed, append the rest of the other list
        current.next = list1 if list1 else list2
        # The sentinel node's next pointer points to the head of the merged list
        return sentinel.next
