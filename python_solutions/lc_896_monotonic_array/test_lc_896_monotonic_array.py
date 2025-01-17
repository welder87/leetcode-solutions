import pytest

from lc_896_monotonic_array import Solution


test_cases = (
    ([1, 2, 2, 3], True),
    ([6, 5, 4, 4], True),
    ([1, 3, 2], False),
    ([-25, -25, -21, -17, -17, 0, 0, 1, 3, 3, 5, 12, 12], True),
    ([-25, -25, -21, -17, -17, 0, 0, 1, 3, 3, 5, 12, 11], False),
    ([12, 12, 5, 3, 3, 1, 0, 0, -17, -17, -21, -25, -25], True),
    ([12, 12, 5, 3, 3, 1, 0, 0, -17, -17, -21, -25, -24], False),
    ([-1, 10, -5], False),
    ([-3, -4, -5, 0], False),
    ([-1, -2, -5], True),
    ([0, 0, 0], True),
    ([-1, -1, -1], True),
    ([1, 1, 1], True),
    ([], False),
    ([1], True),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: bool, solution: Solution):
    assert ans is solution.isMonotonic(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
