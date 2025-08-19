import pytest

from lc_912_sort_an_array.lc_912_sort_an_array import Solution

test_cases = (
    # preset cases
    ([5, 2, 3, 1], [1, 2, 3, 5]),
    ([5, 1, 1, 2, 0, 0], [0, 0, 1, 1, 2, 5]),
    # common cases
    ([-5, -1, 0, -1, 4, -12, 4, -1], [-12, -5, -1, -1, -1, 0, 4, 4]),
    ([-12, -5, -1, -1, -1, 0, 4, 4], [-12, -5, -1, -1, -1, 0, 4, 4]),
    ([4, 4, 0, -1, -1, -1, -5, -12], [-12, -5, -1, -1, -1, 0, 4, 4]),
    ([-1, -5, -1, 0, -1, 4, -12, 4, -1], [-12, -5, -1, -1, -1, -1, 0, 4, 4]),
    # corner cases
    ([-1, -1, -1], [-1, -1, -1]),
    ([], []),
    ([0], [0]),
    ([1, -1], [-1, 1]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert solution.sortArray(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
