import pytest

from lc_456_132_pattern.lc_456 import Solution

test_cases = (
    # preset cases
    ([1, 2, 3, 4], False),
    ([3, 1, 4, 2], True),
    ([-1, 3, 2, 0], True),
    ([-2, 1, 2, -2, 1, 2], True),
    # common cases
    ([1, 2, 3, 4, 5, 6, 7], False),
    ([0, 1, -5, 0, 3, 5, -9, 2], True),
    ([3, -1, 4, 5, 6, 7, 3], True),
    ([3, -1, 4, 5, 8, 8, 3], True),
    ([3, 4, 8, 8, 8, 3], False),
    ([3, 4, 8, 8, 8, 3, 5, 4], True),
    # corner cases
    ([1], False),
    ([1, 1], False),
    ([-1, 3, 0], True),
    ([0, 0, 0], False),
    ([-1, 0, 1], False),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: bool, solution: Solution):
    assert solution.find132pattern(nums) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
