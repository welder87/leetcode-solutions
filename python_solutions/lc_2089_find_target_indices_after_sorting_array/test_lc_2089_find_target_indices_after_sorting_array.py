import pytest

from lc_2089_find_target_indices_after_sorting_array import Solution


test_cases = (
    # preset cases
    ([1, 2, 5, 2, 3], 2, [1, 2]),
    ([1, 2, 5, 2, 3], 3, [3]),
    ([1, 2, 5, 2, 3], 5, [4]),
    # common cases
    ([1, 2, 5, 2, 3, 2], 2, [1, 2, 3]),
    ([1, 2, 5, 5, 3, 5], 5, [3, 4, 5]),
    ([1, 2, 1, 1, 3, 5], 1, [0, 1, 2]),
    ([1, 2, 1, 1, 3, 5], 4, []),
    # corner cases
    ([2, 2, 2], 2, [0, 1, 2]),
    ([2, 2, 2], 3, []),
    ([2], 2, [0]),
    ([1], 2, []),
)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v0(nums: list[int], target: int, ans: int, solution: Solution):
    assert ans == solution.targetIndices(nums, target)


@pytest.mark.parametrize(("nums", "target", "ans"), test_cases)
def test_success_v1(nums: list[int], target: int, ans: int, solution: Solution):
    assert ans == solution.targetIndicesV1(nums, target)


@pytest.fixture
def solution() -> Solution:
    return Solution()
