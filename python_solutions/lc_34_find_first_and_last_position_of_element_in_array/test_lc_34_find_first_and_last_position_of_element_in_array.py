import pytest

from lc_34_find_first_and_last_position_of_element_in_array import Solution


test_cases = (
    # preset cases
    ([5, 7, 7, 8, 8, 10], 8, [3, 4]),
    ([5, 7, 7, 8, 8, 10], 6, [-1, -1]),
    ([], 0, [-1, -1]),
    # common cases
    ([-10, -10, -3, -3, -3, 0, 0, 5, 8, 8, 8, 8], -10, [0, 1]),
    ([-10, -10, -3, -3, -3, 0, 0, 5, 8, 8, 8, 8], -3, [2, 4]),
    ([-10, -10, -3, -3, -3, 0, 0, 5, 8, 8, 8, 8], 0, [5, 6]),
    ([-10, -10, -3, -3, -3, 0, 0, 5, 8, 8, 8, 8], 5, [7, 7]),
    ([-10, -10, -3, -3, -3, 0, 0, 5, 8, 8, 8, 8], 8, [8, 11]),
    # corner cases
    ([-10, -10, -10], -10, [0, 2]),
    ([-10, -10, -10], -11, [-1, -1]),
    ([1], 1, [0, 0]),
    ([1], 0, [-1, -1]),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: list[int], solution: Solution):
    assert solution.searchRange(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: list[int], solution: Solution):
    assert solution.searchRangeV1(nums, target) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
