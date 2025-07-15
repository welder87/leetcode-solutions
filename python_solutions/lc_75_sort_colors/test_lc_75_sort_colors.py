import pytest

from lc_75_sort_colors.lc_75_sort_colors import Solution

test_cases = (
    # preset cases
    ([2, 0, 2, 1, 1, 0], [0, 0, 1, 1, 2, 2]),
    ([2, 0, 1], [0, 1, 2]),
    # common cases
    ([2, 2, 1, 0], [0, 1, 2, 2]),
    ([0, 1, 2, 2], [0, 1, 2, 2]),
    ([0, 2, 1], [0, 1, 2]),
    ([1, 2, 0], [0, 1, 2]),
    # corner cases
    ([1, 1, 1], [1, 1, 1]),
    ([2], [2]),
    ([0], [0]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    solution.sortColors(nums)
    assert nums == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    solution.sortColorsV1(nums)
    assert nums == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: int, solution: Solution):
    solution.sortColorsV2(nums)
    assert nums == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
