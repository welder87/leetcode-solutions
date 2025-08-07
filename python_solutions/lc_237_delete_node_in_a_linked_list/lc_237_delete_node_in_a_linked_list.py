from common_datastructures import ListNode


class Solution:
    def deleteNode(self, node: ListNode) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        nxt = node.next
        while nxt is not None:
            node.val = nxt.val
            if nxt.next is None:
                node.next = None
                break
            node = nxt
            nxt = nxt.next
