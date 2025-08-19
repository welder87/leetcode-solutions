import pytest

from common_datastructures import ListNode, compare_singly_linked_lists
from lc_142_linked_list_cycle_2.lc_142_linked_list_cycle_2 import Solution

_ll0 = ListNode(val=3, next=ListNode(val=2, next=ListNode(val=0)))
_ll0.next.next.next = ListNode(val=-4, next=_ll0.next)
_ll1 = ListNode(val=1)
_ll1.next = ListNode(val=2, next=_ll1)
_ll3 = ListNode(val=3, next=ListNode(val=2, next=ListNode(val=0)))
_ll3.next.next.next = ListNode(val=-4, next=_ll3)
_ll4 = ListNode(
    val=3,
    next=ListNode(val=-2, next=ListNode(val=0, next=ListNode(val=-76))),
)
_ll4.next.next.next.next = ListNode(val=11, next=_ll4.next.next)
_ll5 = ListNode(
    val=3,
    next=ListNode(
        val=-2,
        next=ListNode(val=0, next=ListNode(val=-76, next=ListNode(val=11))),
    ),
)
_ll7 = ListNode(val=3)
_ll7.next = _ll7

test_cases = (
    # preset cases
    # Node(3) → Node(2) → Node(0) → Node(-4)
    #           ↑___________________|
    pytest.param(_ll0, _ll0.next, 3, id="0"),
    # Node(1) → Node(2)
    # ↑_________|
    pytest.param(_ll1, _ll1, 1, id="1"),
    # Node(1)
    pytest.param(ListNode(val=1), None, None, id="2"),
    # common cases
    # Node(3) → Node(2) → Node(0) → Node(-4)
    # ↑_____________________________|
    pytest.param(_ll3, _ll3, 4, id="3"),
    # Node(3) → Node(-2) → Node(0) → Node(-76) → Node(11)
    #                      ↑_____________________|
    pytest.param(_ll4, _ll4.next.next, 3, id="4"),
    # Node(3) → Node(-2) → Node(0) → Node(-76) → Node(11)
    pytest.param(_ll5, None, None, id="5"),
    # corner cases
    # None
    pytest.param(None, None, None, id="6"),
    # Node(3)
    # ↑_____|
    pytest.param(_ll7, _ll7, 1, id="7"),
)


@pytest.mark.parametrize(("lst", "ans", "cycle_count"), test_cases)
def test_success_v0(
    lst: ListNode | None, ans: ListNode | None, cycle_count: int, solution: Solution
):
    assert compare_singly_linked_lists(
        solution.detectCycle(lst),
        ans,
        cycle_count=cycle_count,
    )


@pytest.fixture
def solution() -> Solution:
    return Solution()
