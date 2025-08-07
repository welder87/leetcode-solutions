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

    def deleteNodeV1(self, node):
        # Solution https://leetcode.com/problems/delete-node-in-a-linked-list/editorial/
        # Overwrite data of next node on current node.
        node.val = node.next.val
        # Make current node point to next of next node.
        node.next = node.next.next
