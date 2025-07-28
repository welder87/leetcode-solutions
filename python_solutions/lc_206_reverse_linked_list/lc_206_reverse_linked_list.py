from common_datastructures import ListNode


class Solution:
    def reverseList(self, head: ListNode | None) -> ListNode | None:
        # Time complexity: O(2*n). Space complexity: O(n).
        lst = []
        cur = head
        while cur is not None:
            lst.append(cur.val)
            cur = cur.next
        cur = head
        while cur is not None:
            cur.val = lst.pop()
            cur = cur.next
        return head

    def reverseListV1(self, head: ListNode | None) -> ListNode | None:
        # Time complexity: O(n). Space complexity: O(1).
        prev = None
        while head is not None:
            nxt = head.next
            head.next = prev
            prev = head
            head = nxt
        return prev
