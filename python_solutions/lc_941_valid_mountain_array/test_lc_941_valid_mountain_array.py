import pytest

from lc_941_valid_mountain_array import Solution


test_cases: tuple[list[int], bool] = (
    # preset cases
    ([0, 2, 3, 4, 5, 2, 1, 0], True),
    ([0, 2, 3, 3, 5, 2, 1, 0], False),
    ([2, 1], False),
    ([3, 5, 5], False),
    ([0, 3, 2, 1], True),
    # common cases
    ([0, 2, 3, 4, 5, 2, 1, 1], False),
    ([0, 0, 3, 4, 5, 2, 1, 0], False),
    ([0, 2, 3, 4, 5, 2, 2, 0], False),
    ([1, 3, 1], True),
    ([1, 2, 3, 3, 2], False),
    # corner cases
    ([1, 1, 1], False),
    ([3, 2, 1], False),
    ([1, 2, 3], False),
    ([], False),
    ([1], False),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: bool, solution: Solution):
    assert ans is solution.validMountainArray(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
