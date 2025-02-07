import pytest

from lc_268_missing_number import Solution

test_cases = (
    # preset cases
    ([3, 0, 1], 2),
    ([0, 1], 2),
    ([9, 6, 4, 2, 3, 5, 7, 0, 1], 8),
    ([1, 8, 4, 6, 2, 0, 9, 7, 5], 3),
    # common cases
    ([2, 0, 1], 3),
    ([3, 0, 1, 5, 4], 2),
    ([3, 0, 4, 5, 2], 1),
    ([2, 3, 4, 5, 1], 0),
    ([2, 3, 0, 5, 1], 4),
    # corner cases
    ([0], 1),
    ([1], 0),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.missingNumber(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert solution.missingNumberV1(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: int, solution: Solution):
    assert solution.missingNumberV2(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
