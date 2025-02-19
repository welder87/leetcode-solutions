import pytest

from lc_2733_neither_minimum_nor_maximum import Solution

test_cases = (
    # preset cases
    ([3, 2, 1, 4], {2, 3}),
    ([1, 2], {-1}),
    ([2, 1, 3], {2}),
    # common cases
    ([3, 50, 1, 100, 25], {50, 25, 3}),
    # corner cases
    ([1], {-1}),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: set[int], solution: Solution):
    assert solution.findNonMinOrMax(nums) in ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: set[int], solution: Solution):
    assert solution.findNonMinOrMaxV1(nums) in ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: set[int], solution: Solution):
    assert solution.findNonMinOrMaxV2(nums) in ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
