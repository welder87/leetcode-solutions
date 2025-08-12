import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_24_swap_nodes_in_pairs.lc_24_swap_nodes_in_pairs import Solution

test_cases = (
    # preset cases
    ([1, 2, 3, 4], [2, 1, 4, 3]),
    ([], []),
    ([1], [1]),
    ([1, 2, 3], [2, 1, 3]),
    # common cases
    ([0, 2, 67, 12, 4], [2, 0, 12, 67, 4]),
    ([0, 2, 67, 12, 4, 34], [2, 0, 12, 67, 34, 4]),
    # corner cases
    ([0, 50], [50, 0]),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.swapPairs(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v1(lst: list[int], ans: list[int], solution: Solution):
    res = solution.swapPairsV1(to_singly_linked_list(lst))
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.fixture
def solution() -> Solution:
    return Solution()
