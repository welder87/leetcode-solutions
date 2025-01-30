import pytest

from lc_27_remove_element import Solution

test_cases = (
    # preset cases
    ([3, 2, 2, 3], 3, 2, {2, 2, None, None}),
    ([0, 1, 2, 2, 3, 0, 4, 2], 2, 5, {0, 1, 4, 0, 3, None, None, None}),
    # common cases
    ([2, 1, 2, 2, 3, 0, 4, 2], 2, 4, {4, 1, 0, 3, None, None, None, None}),
    ([2, 1, 0, 2, 3, 0, 4, 2], 0, 6, {2, 1, 2, 2, 3, 4, None, None}),
    ([0, 1, 0, 2, 3, 0, 4, 1], 2, 7, {0, 1, 0, 1, 3, 0, 4, None}),
    ([0, 1, 0, 2, 3, 0, 4, 1], 5, 8, {0, 1, 0, 2, 3, 0, 4, 1}),
    # corner cases
    ([1], 2, 1, {1}),
    ([2], 2, 0, {None}),
    ([2, 2, 2], 2, 0, {None, None, None}),
)


@pytest.mark.parametrize(("nums", "val", "ans", "expected"), test_cases)
def test_success_v0(
    nums: list[int],
    val: int,
    ans: int,
    expected: set[int],
    solution: Solution,
):
    res = solution.removeElement(nums, val)
    assert ans == res
    assert set(nums) == expected


@pytest.fixture
def solution() -> Solution:
    return Solution()
