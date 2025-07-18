import pytest

from lc_2824_count_pairs_whose_sum_is_less_than_target.lc_2824_count_pairs_whose_sum_is_less_than_target import (
    Solution,
)

test_cases = (
    # preset cases
    ([-1, 1, 2, 3, 1], 2, 3),
    ([-6, 2, 5, -2, -7, -1, 3], -2, 10),
    # common cases
    ([-1, 1, 2, 0, 3, 1], 6, 15),
    ([-1, 1, 2, 0, 3, 1], 2, 6),
    ([-1, 1, 2, 0, 3, 1], -1, 0),
    ([-1, -1, -1, -1], 2, 6),
    ([-1, -1, -1, -1], -2, 0),
    ([-1, -1, -1, -1], -3, 0),
    # corner cases
    ([-1, -50], 3, 1),
    ([-1, -49], -50, 0),
    ([-1], 3, 0),
    ([-1], -2, 0),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.countPairs(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.countPairsV1(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v2(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.countPairsV2(nums, target) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
