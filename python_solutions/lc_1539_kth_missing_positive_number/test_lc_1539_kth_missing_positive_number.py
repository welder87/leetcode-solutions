import pytest

from lc_1539_kth_missing_positive_number.lc_1539_kth_missing_positive_number import (
    Solution,
)

test_cases = (
    # preset cases
    ([2, 3, 4, 7, 11], 5, 9),
    ([1, 2, 3, 4], 2, 6),
    # common cases
    ([2, 3, 4, 7, 11], 1, 1),
    ([2, 3, 4, 7, 11], 2, 5),
    ([10, 12, 13, 14], 1, 1),
    ([10, 12, 13, 14], 2, 2),
    ([10, 12, 13, 14], 14, 18),
    ([10, 12, 13, 14], 11, 15),
    ([10, 12, 13, 14], 10, 11),
    ([10, 12, 13, 14], 9, 9),
    # # corner cases
    ([1], 1, 2),
    ([2], 1, 1),
    ([3], 1, 1),
    ([14], 10, 10),
    ([14], 14, 15),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert ans == solution.findKthPositive(nums, target)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert ans == solution.findKthPositiveV1(nums, target)


@pytest.fixture
def solution() -> Solution:
    return Solution()
