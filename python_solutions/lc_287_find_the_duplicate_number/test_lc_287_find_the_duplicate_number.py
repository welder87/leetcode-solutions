import pytest

from lc_287_find_the_duplicate_number import Solution

test_cases = (
    # preset cases
    ([1, 3, 4, 2, 2], 2),
    ([3, 1, 3, 4, 2], 3),
    ([3, 3, 3, 3, 3], 3),
    # common cases
    ([3, 3, 1, 2, 4], 3),
    ([3, 3, 3, 3, 1], 3),
    ([3, 2, 3, 3, 1], 3),
    ([3, 2, 3, 3, 1], 3),
    # corner cases
    ([1, 1], 1),
    ([1], 0),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.findDuplicate(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert solution.findDuplicateV1(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: int, solution: Solution):
    assert solution.findDuplicateV2(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
