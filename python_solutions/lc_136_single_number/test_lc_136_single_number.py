import pytest

from lc_136_single_number import Solution

test_cases = (
    # preset cases
    ([2, 2, 1], 1),
    ([4, 1, 2, 1, 2], 4),
    ([1], 1),
    # common cases
    ([1, -31, 17, -31, 0, 1, 0], 17),
    ([1, -31, 17, 0, 1, 0, 17], -31),
    ([1, 2, 1, 2, 4], 4),
    ([1, 3, 2, 3, 1], 2),
    ([2, 3, 2, 3, 1], 1),
    # corner cases
    ([1, 1, -25], -25),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: int, solution: Solution):
    assert solution.singleNumber(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: int, solution: Solution):
    assert solution.singleNumberV1(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: int, solution: Solution):
    assert solution.singleNumberV2(nums) == ans


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v3(nums: list[int], ans: int, solution: Solution):
    assert solution.singleNumberV3(nums) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
