import pytest

from lc_80_remove_duplicates_from_sorted_array_2.lc_80 import Solution

test_cases = (
    # preset cases
    ([1, 1, 1, 2, 2, 3], 5, [1, 1, 2, 2, 3, None]),
    ([0, 0, 1, 1, 1, 1, 2, 3, 3], 7, [0, 0, 1, 1, 2, 3, 3, None, None]),
    # common cases
    ([-1, -1, -1, -1, 0, 0, 0, 3, 3, 4], 7, [-1, -1, 0, 0, 3, 3, 4, None, None, None]),
    (
        [-1, -1, -1, -1, 0, 0, 0, 3, 3, 4, 4, 4, 4],
        8,
        [-1, -1, 0, 0, 3, 3, 4, 4, None, None, None, None, None],
    ),
    ([-1, 2, 2, 2, 2, 5], 4, [-1, 2, 2, 5, None, None]),
    ([-1, -1, -1, -1, -1], 2, [-1, -1, None, None, None]),
    ([-1, -1, 0, 0, 3, 3], 6, [-1, -1, 0, 0, 3, 3]),
    # corner cases
    ([-15], 1, [-15]),
    ([-15, -15], 2, [-15, -15]),
    ([0, 0, 0], 2, [0, 0, None]),
    ([-15, 0, 1, 4], 4, [-15, 0, 1, 4]),
)


@pytest.mark.parametrize(("nums", "ans", "deduplicated_nums"), test_cases)
def test_success_v0(
    nums: list[int],
    ans: int,
    deduplicated_nums: list[int],
    solution: Solution,
):
    assert solution.removeDuplicates(nums) == ans
    assert nums == deduplicated_nums


@pytest.fixture
def solution() -> Solution:
    return Solution()
