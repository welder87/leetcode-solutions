import pytest

from lc_2460_apply_operations_to_an_array import Solution


test_cases = (
    # preset cases
    ([1, 2, 2, 1, 1, 0], [1, 4, 2, 0, 0, 0]),
    ([0, 1], [1, 0]),
    (
        [847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272],
        [1694, 399, 832, 1758, 412, 206, 272, 0, 0, 0, 0, 0, 0, 0],
    ),
    # common cases
    ([2, 2, 2, 2, 3, 1], [4, 4, 3, 1, 0, 0]),
    ([1, 3, 1, 3, 4, 0], [1, 3, 1, 3, 4, 0]),
    ([0, 3, 1, 3, 4, 2], [3, 1, 3, 4, 2, 0]),
    # corner cases
    ([0, 0], [0, 0]),
    ([2, 2], [4, 0]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.applyOperations(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
