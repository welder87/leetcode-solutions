import pytest

from lc_26_remove_duplicates_from_sorted_array import Solution

test_cases = (
    # preset cases
    ([1, 1, 2], [1, 2, None], 2),
    (
        [0, 0, 1, 1, 1, 2, 2, 3, 3, 4],
        [0, 1, 2, 3, 4, None, None, None, None, None],
        5,
    ),
    # common cases
    (
        [-5, -1, -1, 0, 0, 1, 2, 2, 4, 4],
        [-5, -1, 0, 1, 2, 4, None, None, None, None],
        6,
    ),
    ([1, 2, 2], [1, 2, None], 2),
    ([1, 1, 1, 1, 2], [1, 2, None, None, None], 2),
    # corner cases
    ([-1, -1, -1], [-1, None, None], 1),
    ([0], [0], 1),
)


@pytest.mark.parametrize(("nums", "arr", "ans"), test_cases)
def test_success_v0(
    nums: list[int],
    arr: list[int | None],
    ans: int,
    solution: Solution,
):
    assert ans == solution.removeDuplicates(nums)
    assert nums == arr


@pytest.mark.parametrize(("nums", "arr", "ans"), test_cases)
def test_success_v1(
    nums: list[int],
    arr: list[int | None],
    ans: int,
    solution: Solution,
):
    assert ans == solution.removeDuplicatesV1(nums)
    assert nums == arr


@pytest.fixture
def solution() -> Solution:
    return Solution()
