import pytest

from lc_88_merge_sorted_array import Solution

test_cases = (
    # preset cases
    ([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3, [1, 2, 2, 3, 5, 6]),
    ([1], 1, [], 0, [1]),
    ([0], 0, [1], 1, [1]),
    # common cases
    ([-12, -1, 0, 4, 0, 0, 0], 4, [-2, 0, 5], 3, [-12, -2, -1, 0, 0, 4, 5]),
    ([1, 1, 1, 0, 0, 0], 3, [1, 1, 2], 3, [1, 1, 1, 1, 1, 2]),
    ([-10, -3, -1, 0, 0, 0, 0], 3, [-12, -7, 0, 1], 4, [-12, -10, -7, -3, -1, 0, 1]),
    # corner cases
    ([1, 2, 3, 0, 0, 0], 3, [4, 5, 6], 3, [1, 2, 3, 4, 5, 6]),
    ([4, 5, 6, 0, 0, 0], 3, [1, 2, 3], 3, [1, 2, 3, 4, 5, 6]),
)


@pytest.mark.parametrize(("nums1", "m", "nums2", "n", "ans"), test_cases)
def test_success_v0(
    nums1: list[int],
    m: int,
    nums2: list[int],
    n: int,
    ans: list[int],
    solution: Solution,
):
    solution.merge(nums1, m, nums2, n)
    assert ans == nums1


@pytest.mark.parametrize(("nums1", "m", "nums2", "n", "ans"), test_cases)
def test_success_v1(
    nums1: list[int],
    m: int,
    nums2: list[int],
    n: int,
    ans: list[int],
    solution: Solution,
):
    solution.merge_v1(nums1, m, nums2, n)
    assert ans == nums1


@pytest.fixture
def solution() -> Solution:
    return Solution()
