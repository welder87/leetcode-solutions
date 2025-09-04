import pytest

from lc_496_next_greater_element_1.lc_496_next_greater_element_1 import Solution

test_cases = (
    # preset cases
    # ([4, 1, 2], [1, 3, 4, 2], [-1, 3, -1]),
    # ([2, 4], [1, 2, 3, 4], [3, -1]),
    # # common cases
    # ([9, 0, 24, 17], [11, 9, 4, 0, 20, 24, 17, 35], [20, 20, 35, 35]),
    ([9, 0, 24, 35], [11, 35, 4, 0, 20, 24, 17, 9], [-1, 20, -1, -1]),
    ([9, 0, 24, 35, 11], [11, 35, 4, 9, 0, 20, 17, 24], [20, 20, -1, -1, 35]),
    # corner cases
    ([1], [1], [-1]),
    ([2, 4], [2, 4], [4, -1]),
)


@pytest.mark.parametrize(("nums1", "nums2", "ans"), test_cases)
def test_success_v0(
    nums1: list[int],
    nums2: list[int],
    ans: list[int],
    solution: Solution,
):
    res = solution.nextGreaterElement(nums1, nums2)
    assert res == ans


@pytest.mark.parametrize(("nums1", "nums2", "ans"), test_cases)
def test_success_v1(
    nums1: list[int],
    nums2: list[int],
    ans: list[int],
    solution: Solution,
):
    res = solution.nextGreaterElementV1(nums1, nums2)
    assert res == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
