import pytest

from common_datastructures import (
    compare_singly_linked_lists,
    to_list,
    to_singly_linked_list,
)
from lc_203_remove_linked_list_elements.lc_203_remove_linked_list_elements import (
    Solution,
)

test_cases = (
    # preset cases
    ([1, 2, 6, 3, 4, 5, 6], 6, [1, 2, 3, 4, 5]),
    ([], 1, []),
    ([7, 7, 7, 7], 7, []),
    # common cases
    ([7, 7, 7, 7, 45], 7, [45]),
    ([7, 7, 7, 7, 45, 15], 7, [45, 15]),
    ([7, 7, 50, 7, 7, 45, 15, 7, 7], 7, [50, 45, 15]),
    ([50, 7, 7, 45, 15], 7, [50, 45, 15]),
    ([50, 45, 15], 7, [50, 45, 15]),
    ([45, 7, 7, 7, 7], 7, [45]),
    # corner cases
    ([0], 0, []),
    ([1], 50, [1]),
)


@pytest.mark.parametrize(("head", "val", "ans"), test_cases)
def test_success_v0(head: list[int], val: int, ans: list[int], solution: Solution):
    res = solution.removeElements(to_singly_linked_list(head), val)
    assert ans == to_list(res)
    assert compare_singly_linked_lists(to_singly_linked_list(ans), res)


@pytest.fixture
def solution() -> Solution:
    return Solution()
