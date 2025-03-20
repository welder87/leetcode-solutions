import pytest

from lc_153_find_minimum_in_rotated_sorted_array import Solution


test_cases = (
    # preset cases
    ([3, 4, 5, 1, 2], 1),
    ([4, 5, 6, 7, 0, 1, 2], 0),
    ([11, 13, 15, 17], 11),
    # common cases
    ([4, 6, 7, 8, 9, 0, 1, 2], 0),
    ([17, 11, 13, 15], 11),
    ([2, 3, 4, 5, 1], 1),
    ([4, 5, 1, 2, 3], 1),
    ([5, 1, 2, 3, 4], 1),
    ([5, 8, -7, -3, 0], -7),
    # corner cases
    ([1, 0], 0),
    ([3], 3),
    ([-3], -3),
    ([0], 0),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.findMin(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
