import pytest

from common_datastructures import compare_singly_linked_lists, to_singly_linked_list
from lc_21_merge_two_sorted_lists.lc_21_merge_two_sorted_lists import Solution

test_cases = (
    # preset cases
    ([1, 2, 4], [1, 3, 4], [1, 1, 2, 3, 4, 4]),
    ([], [], []),
    ([], [0], [0]),
    # common cases
    ([1, 2, 3], [2, 5, 6], [1, 2, 2, 3, 5, 6]),
    ([-12, -1, 0, 4], [-2, 0, 5], [-12, -2, -1, 0, 0, 4, 5]),
    ([1, 1, 1], [1, 1, 2], [1, 1, 1, 1, 1, 2]),
    ([-10, -3, -1], [-12, -7, 0, 1], [-12, -10, -7, -3, -1, 0, 1]),
    # corner cases
    ([1, 2, 3], [4, 5, 6], [1, 2, 3, 4, 5, 6]),
    ([4, 5, 6], [1, 2, 3], [1, 2, 3, 4, 5, 6]),
    ([-1], [], [-1]),
)


@pytest.mark.parametrize(("l1", "l2", "ans"), test_cases)
def test_success_v0(l1: list[int], l2: list[int], ans: list[int], solution: Solution):
    slt1 = to_singly_linked_list(l1)
    slt2 = to_singly_linked_list(l2)
    slt = to_singly_linked_list(ans)
    res = solution.mergeTwoLists(slt1, slt2)
    assert compare_singly_linked_lists(slt, res)


@pytest.fixture
def solution() -> Solution:
    return Solution()
