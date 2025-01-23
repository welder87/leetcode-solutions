import pytest

from lc_2441_largest_pos_int_that_exists_with_its_neg import Solution


test_cases = (
    # preset cases
    ([-1, 2, -3, 3], 3),
    ([-1, 10, 6, 7, -7, 1], 7),
    ([-10, 8, 6, 7, -2, -3], -1),
    ([-9, -43, 24, -23, -16, -30, -38, -30], -1),
    ([-25, 25, -27, 45, 31, 46, 46, 21], 25),
    ([-37, 37, -9, 2, 47, 18, 13, -11, 9, -28], 37),
    # common cases
    # corner cases
    ([1], -1),
    ([1, -1], 1),
    ([1, -2], -1),
    ([1, 3], -1),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.findMaxK(nums)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.findMaxKV1(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
