import pytest

from lc_697_degree_of_an_array import Solution

test_cases = (
    # preset cases
    ([1, 2, 2, 3, 1], 2),
    ([1, 2, 2, 3, 1, 4, 2], 6),
    # common cases
    ([2, 1, 2, 3, 1, 4, 2], 7),
    ([0, 3, 1, 2, 2, 3, 1], 2),
    ([1, 1, 2, 2, 3, 3], 2),
    # corner cases
    ([1], 1),
    ([1, 0], 1),
    ([1, 1], 2),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.findShortestSubArray(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
