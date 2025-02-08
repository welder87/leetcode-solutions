import pytest

from lc_645_set_mismatch import Solution

test_cases = (
    # preset cases
    ([1, 2, 2, 4], [2, 3]),
    ([1, 1], [1, 2]),
    # common cases
    ([1, 2, 3, 3], [3, 4]),
    ([4, 2, 1, 4], [4, 3]),
    ([1, 2, 8, 4, 5, 6, 7, 8], [8, 3]),
    ([1, 2, 3, 4, 5, 6, 8, 8], [8, 7]),
    ([8, 2, 4, 3, 7, 6, 5, 8], [8, 1]),
    ([3, 2, 2], [2, 1]),
    ([1, 7, 8, 4, 5, 6, 2, 8], [8, 3]),
    ([3, 3, 2], [3, 1]),
    ([3, 2, 1, 2], [2, 4]),
    # corner cases
    ([2, 2], [2, 1]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert solution.findErrorNums(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: list[int], solution: Solution):
    assert solution.findErrorNumsV1(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: list[int], solution: Solution):
    assert solution.findErrorNumsV2(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
