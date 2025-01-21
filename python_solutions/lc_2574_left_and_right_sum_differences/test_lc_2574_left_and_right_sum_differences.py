import pytest

from lc_2574_left_and_right_sum_differences import Solution


test_cases = (
    # preset cases
    ([10, 4, 8, 3], [15, 1, 11, 22]),
    ([1], [0]),
    # common cases
    ([100, 10, 1, 20], [31, 79, 90, 111]),
    # corner cases
    ([0, 522], [522, 0]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.leftRightDifference(nums)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.leftRightDifferenceV1(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
