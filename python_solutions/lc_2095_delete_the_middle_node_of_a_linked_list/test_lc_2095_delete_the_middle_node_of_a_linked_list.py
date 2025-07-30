import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_2095_delete_the_middle_node_of_a_linked_list.lc_2095_delete_the_middle_node_of_a_linked_list import (
    Solution,
)

test_cases = (
    # preset cases
    pytest.param([1, 3, 4, 7, 1, 2, 6], [1, 3, 4, 1, 2, 6], id="1"),
    pytest.param([1, 2, 3, 4], [1, 2, 4], id="2"),
    pytest.param([2, 1], [2], id="3"),
    # common cases
    pytest.param([1, 99, 15, 45, 4], [1, 99, 45, 4], id="100"),
    pytest.param([1, 99, 15, 45, 4, 4], [1, 99, 15, 4, 4], id="101"),
    # corner cases
    pytest.param([1], [], id="200"),
    pytest.param([1, 2, 1], [1, 1], id="201"),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.deleteMiddle(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.fixture
def solution() -> Solution:
    return Solution()
