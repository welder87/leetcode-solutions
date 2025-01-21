import pytest

from lc_1732_find_the_highest_altitude import Solution


test_cases = (
    # preset cases
    ([-5, 1, 5, 0, -7], 1),
    ([-4, -3, -2, -1, 4, 3, 2], 0),
    # common cases
    # corner cases
    ([0], 0),
    ([1], 1),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.largestAltitude(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
