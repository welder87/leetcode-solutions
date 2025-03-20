import pytest

from lc_33_search_in_rotated_sorted_array import Solution


test_cases = (
    # preset cases
    ([4, 5, 6, 7, 0, 1, 2], 0, 4),
    ([4, 5, 6, 7, 0, 1, 2], 3, -1),
    # common cases
    ([4, 5, 6, 7, 0, 1, 2], -5, -1),
    ([4, 5, 6, 7, 0, 1, 2], 9, -1),
    ([2, 4, 5, 6, 7, 0, 1], 0, 5),
    ([1, 2, 4, 5, 6, 7, 0], 0, 6),
    ([1, 2, 4, 5, 6, 7, 0], 2, 1),
    ([1, 2, 4, 5, 6, 7, 0], 1, 0),
    ([1, 2, 4, 5, 6, 7, 0], 0, 6),
    ([1, 2, 4, 5, 6, 7, 0], 6, 4),
    ([1, 2, 4, 5, 6, 7, 0], 3, -1),
    ([0, 1, 2, 4, 5, 6, 7], 0, 0),
    ([0, 1, 2, 4, 5, 6, 7], 1, 1),
    ([0, 1, 2, 4, 5, 6, 7], 2, 2),
    ([0, 1, 2, 4, 5, 6, 7], 6, 5),
    ([0, 1, 2, 4, 5, 6, 7], 3, -1),
    # corner cases
    ([-1], 3, -1),
    ([-1], -1, 0),
    ([0], 3, -1),
    ([0], 0, 0),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.search(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchV1(nums, target) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
