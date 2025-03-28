import pytest

from lc_1346_check_if_n_and_its_double_exist import Solution


test_cases = (
    # preset cases
    ([10, 2, 5, 3], True),
    ([3, 1, 7, 11], False),
    # common cases
    ([1, -5, -10, 12, 0, 6], True),
    ([-10, -1, 0, 15, 1, 2], True),
    ([-10, -1, 0, 15, 1, 3], False),
    ([-10, -1, 0, 0, 15, 1, 3], True),
    # corner cases
    ([-4, -2], True),
    ([-4, 2], False),
    ([2, 4], True),
    ([-2, 4], False),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: bool, solution: Solution):
    assert solution.checkIfExist(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
