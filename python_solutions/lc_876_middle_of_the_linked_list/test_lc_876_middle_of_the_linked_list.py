import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_876_middle_of_the_linked_list.lc_876_middle_of_the_linked_list import Solution

test_cases = (
    # preset cases
    pytest.param([1, 2, 3, 4, 5], [3, 4, 5], id="1"),
    pytest.param([1, 2, 3, 4, 5, 6], [4, 5, 6], id="2"),
    # common cases
    pytest.param([1, 99, 15, 45, 4], [15, 45, 4], id="100"),
    # corner cases
    pytest.param([1], [1], id="200"),
    pytest.param([1, 1], [1], id="201"),
    pytest.param([1, 1, 1], [1, 1], id="202"),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.middleNode(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.fixture
def solution() -> Solution:
    return Solution()
