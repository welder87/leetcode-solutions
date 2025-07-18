import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_83_remove_duplicates_from_sorted_list.lc_83_remove_duplicates_from_sorted_list import (
    Solution,
)

test_cases = (
    # preset cases
    ([1, 1, 2], [1, 2]),
    ([1, 1, 2, 3, 3], [1, 2, 3]),
    # common cases
    ([-1, 0, 3], [-1, 0, 3]),
    ([-99, -99, 0, 0, 99, 99], [-99, 0, 99]),
    # corner cases
    ([], []),
    ([-1], [-1]),
    ([-1, -1, -1], [-1]),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.deleteDuplicates(to_singly_linked_list(lst))
    assert ans == to_list(res)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), res)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v1(lst: list[int], ans: list[int], solution: Solution):
    res = solution.deleteDuplicatesV1(to_singly_linked_list(lst))
    assert ans == to_list(res)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), res)


@pytest.fixture
def solution() -> Solution:
    return Solution()
