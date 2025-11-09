import pytest

from lc_2149_rearrange_array_elements_by_sign.lc_2149 import Solution

test_cases = (
    # preset cases
    # ([3, 1, -2, -5, 2, -4], [3, -2, 1, -5, 2, -4]),
    # ([-1, 1], [1, -1]),
    # common cases
    ([-1, 1, -4, -5, 3, 4], [1, -1, 3, -4, 4, -5]),
    # corner cases
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert solution.rearrangeArray(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
