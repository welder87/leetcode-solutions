import pytest

from lc_704_binary_search import Solution


test_cases = (
    # preset cases
    ([-1, 0, 3, 5, 9, 12], 9, 4),
    ([-1, 0, 3, 5, 9, 12], 2, -1),
    # common cases
    ([-36, -6, -1, 0, 1, 5, 9, 12], 2, -1),
    ([-36, -6, -1, 0, 1, 5, 9, 12], -7, -1),
    ([-36, -6, -1, 0, 1, 5, 9, 12], -36, 0),
    ([-36, -6, -1, 0, 1, 5, 9, 12], -6, 1),
    ([-36, -6, -1, 0, 1, 5, 9, 12], 0, 3),
    # corner cases
    ([-1], -1, 0),
    ([-1], 0, -1),
    ([-1], 1, -1),
    ([0], 0, 0),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.search(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchV1(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v2(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchV2(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v3(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchV3(nums, target) == ans


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v4(nums: list[int], target: int, ans: int, solution: Solution):
    assert solution.searchV4(nums, target) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
