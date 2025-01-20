import pytest

from lc_922_sort_array_by_parity_2 import Solution

common_result = [[4, 5, 2, 7], [4, 7, 2, 5], [2, 5, 4, 7], [2, 7, 4, 5]]
test_cases = (
    # preset cases
    ([4, 2, 5, 7], common_result),
    ([2, 3], [[2, 3]]),
    # common cases
    ([2, 4, 5, 7], common_result),
    ([2, 4, 7, 5], common_result),
    ([2, 7, 4, 5], common_result),
    ([7, 2, 4, 5], common_result),
    ([7, 4, 2, 5], common_result),
    ([7, 4, 5, 2], common_result),
    ([4, 7, 5, 2], common_result),
    ([4, 5, 7, 2], common_result),
    ([4, 5, 2, 7], common_result),
    ([4, 2, 7, 5], common_result),
    ([4, 7, 2, 5], common_result),
    ([5, 4, 7, 2], common_result),
    ([5, 7, 4, 2], common_result),
    # corner cases
    ([3, 2], [[2, 3]]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[list[int]], solution: Solution):
    assert solution.sortArrayByParityII(nums) in ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: list[list[int]], solution: Solution):
    assert solution.sortArrayByParityIIV1(nums) in ans

@pytest.fixture
def solution() -> Solution:
    return Solution()
