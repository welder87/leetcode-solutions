import pytest

from lc_852_peak_index_in_a_mountain_array import Solution


test_cases = (
    # preset cases
    ([0, 1, 0], 1),
    ([0, 2, 1, 0], 1),
    ([0, 10, 5, 2], 1),
    # common cases
    ([5, 6, 4, 3, 2, 1, 0], 1),
    ([5, 6, 15, 25, 31, 7], 4),
    ([5, 6, 15, 14, 11, 0], 2),
    # corner cases
    ([7, 8, 4], 1),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.peakIndexInMountainArray(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
