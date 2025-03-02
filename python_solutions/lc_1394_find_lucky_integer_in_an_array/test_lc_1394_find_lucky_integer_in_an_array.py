import pytest

from lc_1394_find_lucky_integer_in_an_array import Solution

test_cases = (
    # preset cases
    ([2, 2, 3, 4], 2),
    ([1, 2, 2, 3, 3, 3], 3),
    ([2, 2, 2, 3, 3], -1),
    # common cases
    ([1, 2, 2, 2, 3, 3, 3], 3),
    ([1, 4, 4, 4, 4, 3, 3, 3], 4),
    # corner cases
    ([1], 1),
    ([2], -1),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.findLucky(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert solution.findLuckyV1(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: int, solution: Solution):
    assert solution.findLuckyV2(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
