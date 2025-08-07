import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_237_delete_node_in_a_linked_list.lc_237_delete_node_in_a_linked_list import (
    Solution,
)

test_cases = (
    # preset cases
    pytest.param([4, 5, 1, 9], 5, [4, 1, 9], id="1"),
    pytest.param([4, 5, 1, 9], 1, [4, 5, 9], id="2"),
    # common cases
    pytest.param([-1, 0, 2, 3, -25, 5], -1, [0, 2, 3, -25, 5], id="100"),
    pytest.param([-1, 0, 2, 3, -25, 5], 0, [-1, 2, 3, -25, 5], id="101"),
    pytest.param([-1, 0, 2, 3, -25, 5], 2, [-1, 0, 3, -25, 5], id="102"),
    pytest.param([-1, 0, 2, 3, -25, 5], 3, [-1, 0, 2, -25, 5], id="102"),
    pytest.param([-1, 0, 2, 3, -25, 5], -25, [-1, 0, 2, 3, 5], id="102"),
    # corner cases
    pytest.param([-1, 0], -1, [0], id="200"),
    pytest.param([-1, 0, 1], -1, [0, 1], id="201"),
    pytest.param([-1, 0, 1], 0, [-1, 1], id="202"),
)


@pytest.mark.parametrize(("lst", "val", "ans"), test_cases)
def test_success_v0(lst: list[int], val: int, ans: list[int], solution: Solution):
    ll = to_singly_linked_list(lst)
    node = ll
    while node.next is not None and node.val != val:
        node = node.next
    solution.deleteNode(node)
    assert ans == to_list(ll)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), ll)


@pytest.fixture
def solution() -> Solution:
    return Solution()
