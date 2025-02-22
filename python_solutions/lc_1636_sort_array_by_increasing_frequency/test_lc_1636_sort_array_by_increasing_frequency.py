import pytest

from lc_1636_sort_array_by_increasing_frequency import Solution


test_cases = (
    # preset cases
    ([1, 1, 2, 2, 2, 3], [3, 1, 1, 2, 2, 2]),
    ([2, 3, 1, 3, 2], [1, 3, 3, 2, 2]),
    ([-1, 1, -6, 4, 5, -6, 1, 4, 1], [5, -1, 4, 4, -6, -6, 1, 1, 1]),
    # common cases
    # corner cases
    ([1], [1]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.frequencySort(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
