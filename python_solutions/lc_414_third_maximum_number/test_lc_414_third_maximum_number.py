import pytest

from lc_414_third_maximum_number import Solution

test_cases = (
    # preset cases
    ([3, 2, 1], 1),
    ([1, 2], 2),
    ([2, 2, 3, 1], 1),
    # common cases
    ([1, 2, 3, 4, 1], 2),
    ([10, 1, 3, 4, 1], 3),
    ([2, 2, 2, 4, 1], 1),
    ([-2, -2, 4, -2, 1], -2),
    # corner cases
    ([2], 2),
    ([2, 2], 2),
    ([2, -3], 2),
    ([2, -3, -3], 2),
    ([2, -3, 4], -3),
    ([1, 2, 2], 2),
    ([2, 2, 2], 2),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.thirdMax(nums)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.thirdMaxV1(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
