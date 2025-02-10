import pytest

from lc_1470_shuffle_the_array import Solution

test_cases = (
    # preset cases
    ([2, 5, 1, 3, 4, 7], 3, [2, 3, 5, 4, 1, 7]),
    ([1, 2, 3, 4, 4, 3, 2, 1], 4, [1, 4, 2, 3, 3, 2, 4, 1]),
    ([1, 1, 2, 2], 2, [1, 2, 1, 2]),
    # common cases
    # corner cases
    ([1, 2], 1, [1, 2]),
)


@pytest.mark.parametrize(("nums", "n", "ans"), test_cases)
def test_success_v0(nums: list[int], n: int, ans: list[int], solution: Solution):
    assert solution.shuffle(nums, n) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
