import pytest

from lc_349_intersection_of_two_arrays import Solution

test_cases = (
    # preset cases
    ([1, 2, 2, 1], [2, 2], {2}),
    ([4, 9, 5], [9, 4, 9, 8, 4], {9, 4}),
    # common cases
    ([1, 5, 1, 3, 5, 7], [1, 3], {1, 3}),
    ([1, 2, 3], [4, 6, 5], set()),
    ([4, 6, 5], [5, 6, 4], {4, 6, 5}),
    # corner cases
    ([], [], set()),
)


@pytest.mark.parametrize(("nums1", "nums2", "ans"), test_cases)
def test_success_v0(
    nums1: list[int],
    nums2: list[int],
    ans: set[int],
    solution: Solution,
):
    res = solution.intersection(nums1, nums2)
    assert set(res) == ans


@pytest.mark.parametrize(("nums1", "nums2", "ans"), test_cases)
def test_success_v1(
    nums1: list[int],
    nums2: list[int],
    ans: list[list[int]],
    solution: Solution,
):
    res = solution.intersection_v1(nums1, nums2)
    assert set(res) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
