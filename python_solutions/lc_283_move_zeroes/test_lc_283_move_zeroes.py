import pytest

from lc_283_move_zeroes import Solution

test_cases = (
    # preset cases
    ([0, 1, 0, 3, 12], [1, 3, 12, 0, 0]),
    ([0], [0]),
    # common cases
    ([1, 3, 12, 0, 0], [1, 3, 12, 0, 0]),
    ([0, 0, 0, 1], [1, 0, 0, 0]),
    ([0, 1, 3, 12, 0, 0], [1, 3, 12, 0, 0, 0]),
    ([1, 3, 0, 12, 0], [1, 3, 12, 0, 0]),
    ([12, 3, 0, 1, 0], [12, 3, 1, 0, 0]),
    ([1, 6, 1], [1, 6, 1]),
    # corner cases
    ([0, 0, 0], [0, 0, 0]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: set[int], solution: Solution):
    solution.moveZeroes(nums)
    assert nums == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: set[int], solution: Solution):
    solution.moveZeroesV1(nums)
    assert nums == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
