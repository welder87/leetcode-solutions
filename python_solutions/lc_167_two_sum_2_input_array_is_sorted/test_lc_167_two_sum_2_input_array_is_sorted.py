import pytest

from lc_167_two_sum_2_input_array_is_sorted import Solution

test_cases = (
    # preset cases
    ([2, 7, 11, 15], 9, [1, 2]),
    ([2, 3, 4], 6, [1, 3]),
    ([-1, 0], -1, [1, 2]),
    # common cases
    ([-15, -3, 0, 11, 15], 8, [2, 4]),
    ([-15, -3, 0, 11, 15], 0, [1, 5]),
    ([-15, -3, 0, 11, 15], 11, [3, 4]),
    # corner cases
    ([0, 0], 0, [1, 2]),
    ([-1, -1], -2, [1, 2]),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: list[int], solution: Solution):
    assert ans == solution.twoSum(nums, target)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: list[int], solution: Solution):
    assert ans == solution.twoSumV1(nums, target)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v2(nums: list[int], target: int, ans: list[int], solution: Solution):
    assert ans == solution.twoSumV2(nums, target)


@pytest.fixture
def solution() -> Solution:
    return Solution()
