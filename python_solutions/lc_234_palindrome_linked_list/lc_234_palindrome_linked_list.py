from typing import Optional

from common_datastructures import ListNode


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        # Time complexity: O(n + n/2). Space complexity: O(n).
        if head is None:
            return False
        lst = []
        while head is not None:
            lst.append(head.val)
            spam = head.next
            head.next = None
            head = spam
        i, j = 0, len(lst) - 1
        while i < j:
            if lst[i] != lst[j]:
                return False
            i += 1
            j -= 1
        return True

    def isPalindromeV1(self, head: Optional[ListNode]) -> bool:
        # Time complexity: O(n/2 + 1 + n/2 + n/2 + 1). Space complexity: O(1).
        if head is None:
            return False
        new_head = self._reverse(self._find_mid(head))
        return self._compare(head, new_head)

    def _find_mid(self, head: ListNode) -> ListNode:
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
        return slow

    def _reverse(self, head: ListNode) -> ListNode:
        prev = None
        while head is not None:
            nxt = head.next
            head.next = prev
            prev = head
            head = nxt
        return prev

    def _compare(self, head1: ListNode, head2: ListNode) -> bool:
        while head1 is not None and head2 is not None:
            if head1.val != head2.val:
                return False
            head1 = head1.next
            head2 = head2.next
        return True
