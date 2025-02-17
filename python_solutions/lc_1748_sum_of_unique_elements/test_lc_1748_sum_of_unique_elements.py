import pytest

from lc_1748_sum_of_unique_elements import Solution


test_cases = (
    # preset cases
    ([1, 2, 3, 2], 4),
    ([1, 1, 1, 1, 1], 0),
    ([1, 2, 3, 4, 5], 15),
    # common cases
    ([1, 1, 1, 5], 5),
    ([1, 5, 5, 5], 1),
    # corner cases
    ([1], 1),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert ans == solution.sumOfUnique(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
