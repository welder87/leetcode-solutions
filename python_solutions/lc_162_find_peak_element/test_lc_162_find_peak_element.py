import pytest

from lc_162_find_peak_element import Solution


test_cases = (
    # preset cases
    ([1, 2, 3, 1], {2}),
    ([1, 2, 1, 3, 5, 6, 4], {1, 5}),
    # common cases
    ([7, 6, 4, 5, 6, 7, 8], {0, 6}),
    ([-7, -6, -8, 5, 8, 6, 3, 9, 8], {1, 4, 7}),
    ([7, 8, 4], {1}),
    ([7, 6, 4, 3, 2, -1, -5], {0}),
    # corner cases
    ([7, 6, 4], {0}),
    ([-1], {0}),
    ([-1, 0], {1}),
    ([0, -1], {0}),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.findPeakElement(nums) in ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert solution.findPeakElementV1(nums) in ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
