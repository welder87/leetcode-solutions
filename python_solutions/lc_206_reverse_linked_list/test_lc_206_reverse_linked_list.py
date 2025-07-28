import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_206_reverse_linked_list.lc_206_reverse_linked_list import Solution

test_cases = (
    # preset cases
    pytest.param([1, 2, 3, 4, 5], [5, 4, 3, 2, 1], id="0"),
    pytest.param([1, 2], [2, 1], id="1"),
    pytest.param([], [], id="2"),
    # common cases
    pytest.param([-5, 0, 0, 3, -9, -9, 45], [45, -9, -9, 3, 0, 0, -5], id="3"),
    # corner cases
    pytest.param([1], [1], id="4"),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.reverseList(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v1(lst: list[int], ans: list[int], solution: Solution):
    res = solution.reverseListV1(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.fixture
def solution() -> Solution:
    return Solution()
