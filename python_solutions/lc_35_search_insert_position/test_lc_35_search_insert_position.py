import pytest

from lc_35_search_insert_position import Solution


test_cases = (
    # preset cases
    ([1, 3, 5, 6], 5, 2),
    ([1, 3, 5, 6], 2, 1),
    ([1, 3, 5, 6], 7, 4),
    # common cases
    ([-10, -1, 0, 3, 17], -9, 1),
    ([-10, -3, 0, 3, 17], -1, 2),
    ([-10, -3, 0, 3, 17], 0, 2),
    ([-10, -3, 0, 3, 17], 5, 4),
    # corner cases
    ([-10, -1, 0, 3, 17], -12, 0),
    ([-10, -3, 0, 3, 17], 18, 5),
    ([-1], 1, 1),
    ([-1], -2, 0),
    ([1], 0, 0),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchInsert(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchInsertV1(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v2(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchInsertV2(nums, target) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
