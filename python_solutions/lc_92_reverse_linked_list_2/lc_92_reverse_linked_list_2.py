from typing import Optional

from common_datastructures import ListNode


class Solution:
    def reverseBetween(
        self,
        head: Optional[ListNode],
        left: int,
        right: int,
    ) -> Optional[ListNode]:
        # Time complexity: O(n+m). Space complexity: O(1).
        cur = head
        counter = 0
        pre_start, start, end, post_end, prev = None, None, None, None, None
        while cur is not None:
            counter += 1
            if counter == left:
                start = cur
                pre_start = prev
            if counter == right:
                end = cur
                post_end = cur.next
                break
            prev = cur
            cur = cur.next

        if pre_start is not None:
            pre_start.next = None

        if end is not None:
            end.next = None

        prev, cur = None, start
        while cur is not None:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt

        if pre_start is None and post_end is None:
            return end
        if pre_start is None and post_end is not None:
            start.next = post_end
            return end
        if pre_start is not None and post_end is None:
            pre_start.next = end
            return head

        pre_start.next = end
        start.next = post_end

        return head

    def reverseBetweenV1(
        self,
        head: Optional[ListNode],
        left: int,
        right: int,
    ) -> Optional[ListNode]:
        # Time complexity: O(n). Space complexity: O(1).
        if head is None or head.next is None or left == right:
            return head
        counter = 0
        cur, prev, nxt = head, None, None
        pre_start, start, end, post_end = None, None, None, None
        while cur is not None:
            counter += 1
            if counter < left:
                prev = cur
                cur = cur.next
            elif left <= counter <= right:
                if counter == left:
                    pre_start = prev
                    start = cur
                    if pre_start is not None:
                        pre_start.next = None
                    prev = None
                elif counter == right:
                    post_end = cur.next
                    cur.next = None
                    end = cur
                nxt = cur.next
                cur.next = prev
                prev = cur
                cur = nxt
            else:
                break
        if pre_start is None and post_end is None:
            return end
        if pre_start is None and post_end is not None:
            start.next = post_end
            return end
        if pre_start is not None and post_end is None:
            pre_start.next = end
            return head
        pre_start.next = end
        start.next = post_end
        return head

    def reverseBetweenV2(
        self,
        head: Optional[ListNode],
        left: int,
        right: int,
    ) -> Optional[ListNode]:
        # Solution https://algo.monster/liteproblems/92
        # If the list only contains one node or no reversal is needed, return the head
        # as is.
        if head is None or head.next is None or left == right:
            return head
        # Initialize a dummy node to simplify edge cases
        # where the reversal might include the head of the list.
        dummy = ListNode(0)
        dummy.next = head
        # This node will eventually point to the node right before
        # the reversal starts. Initialize it to the dummy node.
        predecessor = dummy
        # Move the predecessor to the node right before where
        # the reversal is supposed to start.
        for _ in range(left - 1):
            predecessor = predecessor.next
        # Initialize the 'reverse_start' node, which will eventually point to the
        # first node in the sequence that needs to be reversed.
        reverse_start = predecessor.next
        current = reverse_start
        prev = None
        # This loop reverses the nodes between 'left' and 'right'.
        for _ in range(right - left + 1):
            next_temp = current.next
            current.next = prev
            prev = current
            current = next_temp
        # Reconnect the reversed section back to the list.
        predecessor.next = prev  # Connect with node before reversed part.
        reverse_start.next = (
            current  # Connect the last reversed node to the remainder of the list.
        )
        # Return the new head of the list, which is the next of dummy node.
        return dummy.next
