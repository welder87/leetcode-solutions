import pytest

from lc_713_subarray_product_less_than_k import Solution

test_cases = (
    # preset cases
    ([10, 5, 2, 6], 100, 8),
    ([1, 2, 3], 0, 0),
    ([1, 2, 3, 4, 5], 1, 0),
    # common cases
    ([10, 5, 2, 6], 13, 6),
    ([10, 5, 2, 12, 6], 12, 5),
    ([10, 5, 2, 6], 1000, 10),
    # corner cases
    ([1, 1], 1, 0),
    ([1], 1, 0),
    ([1], 2, 1),
    ([1], 0, 0),
)


@pytest.mark.parametrize(("nums", "k", "ans"), test_cases)
def test_success_v0(nums: list[int], k: int, ans: int, solution: Solution):
    assert solution.numSubarrayProductLessThanK(nums, k) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
