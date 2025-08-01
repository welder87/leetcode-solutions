import pytest

from common_datastructures import (
    ListNode,
    compare_singly_linked_lists,
    to_list,
)
from lc_160_intersection_of_two_linked_lists.lc_160_intersection_of_two_linked_lists import (
    Solution,
)

tc1 = ListNode(val=8, next=ListNode(val=4, next=ListNode(val=5)))
ta1 = ListNode(val=4, next=ListNode(val=1, next=tc1))
tb1 = ListNode(val=5, next=ListNode(val=6, next=ListNode(val=1, next=tc1)))

tc2 = ListNode(val=8, next=ListNode(val=4, next=ListNode(val=5)))
ta2 = ListNode(val=4, next=ListNode(val=1, next=tc2))
tb2 = ListNode(val=5, next=ListNode(val=6, next=ListNode(val=1, next=tc2)))

tc3 = None
ta3 = ListNode(val=2, next=ListNode(val=6, next=ListNode(val=4)))
tb3 = ListNode(val=1, next=ListNode(val=5))

tc100 = ListNode(val=8, next=ListNode(val=4, next=ListNode(val=5)))
ta100 = ListNode(val=5, next=ListNode(val=6, next=ListNode(val=1, next=tc100)))
tb100 = ListNode(val=4, next=ListNode(val=1, next=tc100))

tc101 = ListNode(val=15)
ta101 = ListNode(val=2, next=ListNode(val=6, next=ListNode(val=4, next=tc101)))
tb101 = ListNode(val=1, next=ListNode(val=5, next=tc101))

tc200 = None
ta200 = ListNode(val=2)
tb200 = ListNode(val=1)

tc201 = ListNode(val=15, next=ListNode(val=5))
ta201 = ListNode(val=2, next=tc201)
tb201 = ListNode(val=1, next=tc201)

test_cases = (
    # preset cases
    pytest.param(ta1, tb1, tc1, id="1"),
    pytest.param(ta2, tb2, tc2, id="2"),
    pytest.param(ta3, tb3, tc3, id="3"),
    # common cases
    pytest.param(ta100, tb100, tc100, id="100"),
    pytest.param(ta101, tb101, tc101, id="101"),
    # corner cases
    pytest.param(ta200, tb200, tc200, id="200"),
    pytest.param(ta201, tb201, tc201, id="201"),
)


@pytest.mark.parametrize(("lst1", "lst2", "ans"), test_cases)
def test_success_v0(
    lst1: ListNode,
    lst2: ListNode,
    ans: ListNode | None,
    solution: Solution,
):
    res = solution.getIntersectionNode(lst1, lst2)
    assert to_list(res) == to_list(ans)
    assert compare_singly_linked_lists(res, ans)


@pytest.mark.parametrize(("lst1", "lst2", "ans"), test_cases)
def test_success_v1(
    lst1: ListNode,
    lst2: ListNode,
    ans: ListNode | None,
    solution: Solution,
):
    res = solution.getIntersectionNodeV1(lst1, lst2)
    assert to_list(res) == to_list(ans)
    assert compare_singly_linked_lists(res, ans)


@pytest.fixture
def solution() -> Solution:
    return Solution()
