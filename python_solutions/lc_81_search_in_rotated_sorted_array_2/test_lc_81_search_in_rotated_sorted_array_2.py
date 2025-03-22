import pytest

from lc_81_search_in_rotated_sorted_array_2 import Solution


test_cases = (
    # preset cases
    ([2, 5, 6, 0, 0, 1, 2], 0, True),
    ([2, 5, 6, 0, 0, 1, 2], 3, False),
    ([1, 0, 1, 1, 1], 0, True),
    # common cases
    ([2, 5, 6, 0, 0, 0, 0, 1, 2], 0, True),
    ([2, 5, 6, 0, 1, 2], 0, True),
    ([2, 5, 6, -4, -3, 0, 1, 1, 2], 0, True),
    ([2, 5, 6, -4, -3, 0, 1, 1, 2], 1, True),
    ([2, 5, 6, -4, -3, 0, 1, 1, 2], -1, False),
    # corner cases
    ([3, 1, 1], 3, True),
    ([1, 1, 3, 1], 3, True),
    ([1, 1, 0, 1, 1], 0, True),
    ([1, 1, 1, 0, 1], 0, True),
    ([2, 2, 2], 0, False),
    ([-2, -2, -2], -2, True),
    ([-1], 3, False),
    ([-1], -1, True),
    ([0], 3, False),
    ([0], 0, True),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: bool, solution: Solution):
    assert solution.search(nums, target) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
