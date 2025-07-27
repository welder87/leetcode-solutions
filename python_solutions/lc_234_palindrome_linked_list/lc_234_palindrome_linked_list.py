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
