import pytest

from lc_2210_count_hills_and_valleys_in_an_array.lc_2210_count_hills_and_valleys_in_an_array import (
    Solution,
)

test_cases = (
    # preset cases
    ([2, 4, 1, 1, 6, 5], 3),
    ([6, 6, 5, 5, 4, 1], 0),
    # common cases
    # corner cases
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.countHillValley(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
