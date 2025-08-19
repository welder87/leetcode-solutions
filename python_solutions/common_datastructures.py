from typing import Optional


class ListNode:
    """Definition for singly-linked list."""

    def __init__(self, val: int, next: Optional["ListNode"] = None):
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        return f"[ {self.__class__.__name__}: val={self.val}, next={self.next} ]"


def to_list(node: ListNode | None) -> list:
    if node is None:
        return []
    nxt = node
    res = []
    while nxt is not None:
        res.append(nxt.val)
        nxt = nxt.next
    return res


def to_singly_linked_list(lst: list) -> ListNode | None:
    if not lst:
        return None
    head = ListNode(val=lst[0])
    prev = head
    for i in range(1, len(lst)):
        prev.next = ListNode(val=lst[i])
        prev = prev.next
    return head


def compare_singly_linked_lists(
    head1: ListNode | None,
    head2: ListNode | None,
    *,
    cycle_count: int | None = None,
):
    counter = 0
    while head1 and head2:
        counter += 1
        if head1.val != head2.val:
            return False
        head1 = head1.next
        head2 = head2.next
        if counter == cycle_count:
            return True
    return head1 is None and head2 is None
