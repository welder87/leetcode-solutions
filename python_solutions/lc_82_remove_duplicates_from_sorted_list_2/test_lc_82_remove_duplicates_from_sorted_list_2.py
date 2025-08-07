import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_82_remove_duplicates_from_sorted_list_2.lc_82_remove_duplicates_from_sorted_list_2 import (
    Solution,
)

test_cases = (
    # preset cases
    pytest.param([1, 2, 3, 3, 4, 4, 5], [1, 2, 5], id="1"),
    pytest.param([1, 1, 1, 2, 3], [2, 3], id="2"),
    # common cases
    pytest.param([-1, -1, -1, 0, 2, 3, 3, 3, 4, 4, 5, 5, 5], [0, 2], id="100"),
    pytest.param([-1, -1, -1, -1], [], id="101"),
    pytest.param([-1, -1, -1, -1, 0], [0], id="102"),
    pytest.param([-1, -1, -1, -1, 0, 0], [], id="103"),
    pytest.param([-1, -1, -1, -1, 0, 0, 54], [54], id="104"),
    pytest.param([-1, -1, -1, -1, 0, 54, 54], [0], id="105"),
    pytest.param([-1, 0, 3], [-1, 0, 3], id="106"),
    # corner cases
    pytest.param([], [], id="201"),
    pytest.param([-1], [-1], id="202"),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    res = solution.deleteDuplicates(to_singly_linked_list(lst))
    assert ans == to_list(res)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), res)


@pytest.fixture
def solution() -> Solution:
    return Solution()
