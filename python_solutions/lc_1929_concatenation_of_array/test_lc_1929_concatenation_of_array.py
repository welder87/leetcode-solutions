import pytest

from lc_1929_concatenation_of_array import Solution


test_cases = (
    # preset cases
    ([1, 2, 1], [1, 2, 1, 1, 2, 1]),
    ([1, 3, 2, 1], [1, 3, 2, 1, 1, 3, 2, 1]),
    # common cases
    ([1, 10, 5, 3], [1, 10, 5, 3, 1, 10, 5, 3]),
    # corner cases
    ([1], [1, 1]),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.getConcatenation(nums)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.getConcatenationV1(nums)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: list[int], solution: Solution):
    assert ans == solution.getConcatenationV2(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
