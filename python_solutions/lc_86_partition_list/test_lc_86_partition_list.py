import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_86_partition_list.lc_86_partition_list import Solution

test_cases = (
    # preset cases
    ([1, 4, 3, 2, 5, 2], 3, [1, 2, 2, 4, 3, 5]),
    ([2, 1], 2, [1, 2]),
    # common cases
    ([4, 3, 5, -2, 0, -2], 3, [-2, 0, -2, 4, 3, 5]),
    ([4, 3, 5, -2, 0, -2], 5, [4, 3, -2, 0, -2, 5]),
    ([4, 3, 5, -2, 0, -2], 1, [-2, 0, -2, 4, 3, 5]),
    ([4, 3, 5, -2, 0, -2], 0, [-2, -2, 4, 3, 5, 0]),
    ([4, 3, 5, -2, 0, -2], -2, [4, 3, 5, -2, 0, -2]),
    ([4, 3, 5, -2, 0, -2], 10, [4, 3, 5, -2, 0, -2]),
    ([4, 3, 5, -2, 0, -2], -5, [4, 3, 5, -2, 0, -2]),
    # corner cases
    ([-1], 2, [-1]),
    ([5], 2, [5]),
    ([], 2, []),
)


@pytest.mark.parametrize(("lst", "x", "ans"), test_cases)
def test_success_v0(lst: list[int], x: int, ans: list[int], solution: Solution):
    res = solution.partition(to_singly_linked_list(lst), x)
    assert to_list(res) == ans
    assert compare_singly_linked_lists(res, to_singly_linked_list(ans))


@pytest.fixture
def solution() -> Solution:
    return Solution()
