import pytest

from lc_3364_minimum_positive_sum_subarray import Solution


test_cases = (
    # preset cases
    ([3, -2, 1, 4], 2, 3, 1),
    ([-2, 2, -3, 1], 2, 3, -1),
    ([1, 2, 3, 4], 2, 4, 3),
    # common cases
    ([3, -2, 1, 4], 1, 1, 1),
    ([3, -3, 2, 4], 1, 3, 2),
    # corner cases
    ([1], 0, 0, -1),
    ([0], 0, 0, -1),
)


@pytest.mark.parametrize(("nums", "l", "r", "ans"), test_cases)
def test_success_v0(
    nums: list[int],
    l: int,
    r: int,
    ans: list[int],
    solution: Solution,
):
    assert ans == solution.minimumSumSubarray(nums, l, r)


@pytest.fixture
def solution() -> Solution:
    return Solution()
