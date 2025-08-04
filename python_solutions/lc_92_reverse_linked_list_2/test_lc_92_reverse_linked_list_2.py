import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_92_reverse_linked_list_2.lc_92_reverse_linked_list_2 import Solution

test_cases = (
    # preset cases
    pytest.param([1, 2, 3, 4, 5], 2, 4, [1, 4, 3, 2, 5], id="1"),
    pytest.param([5], 1, 1, [5], id="2"),
    # common cases
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        1,
        8,
        [0, -17, 45, 45, 0, -6, 8, 8],
        id="100",
    ),
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        1,
        7,
        [-17, 45, 45, 0, -6, 8, 8, 0],
        id="101",
    ),
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        2,
        8,
        [8, 0, -17, 45, 45, 0, -6, 8],
        id="102",
    ),
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        4,
        6,
        [8, 8, -6, 45, 45, 0, -17, 0],
        id="103",
    ),
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        5,
        5,
        [8, 8, -6, 0, 45, 45, -17, 0],
        id="104",
    ),
    pytest.param(
        [8, 8, -6, 0, 45, 45, -17, 0],
        2,
        7,
        [8, -17, 45, 45, 0, -6, 8, 0],
        id="105",
    ),
    # corner cases
    pytest.param([], 1, 1, [], id="201"),
    pytest.param([1, 2], 1, 2, [2, 1], id="202"),
    pytest.param([1, 2, 3], 1, 2, [2, 1, 3], id="203"),
    pytest.param([1, 2, 3], 1, 3, [3, 2, 1], id="204"),
    pytest.param([1, 2, 3], 2, 3, [1, 3, 2], id="205"),
    pytest.param([1, 2, 3], 3, 3, [1, 2, 3], id="205"),
)


@pytest.mark.parametrize(("lst", "left", "right", "ans"), test_cases)
def test_success_v0(
    lst: list[int],
    left: int,
    right: int,
    ans: list[int],
    solution: Solution,
):
    res = solution.reverseBetween(to_singly_linked_list(lst), left, right)
    assert ans == to_list(res)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), res)


@pytest.fixture
def solution() -> Solution:
    return Solution()
