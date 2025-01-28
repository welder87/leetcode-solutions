import pytest

from lc_1207_unique_number_of_occurrences import Solution


test_cases = (
    # preset cases
    ([1, 2, 2, 1, 1, 3], True),
    ([1, 2], False),
    ([-3, 0, 1, -3, 1, 1, 1, -3, 10, 0], True),
    # common cases
    ([-3, 1, -3, 1, 1, -3], False),
    # corner cases
    ([2, 2], True),
    ([2, 1, 2], True),
    ([1], True),
    ([-1], True),
    ([0], True),
)


@pytest.mark.parametrize(("arr", "ans"), test_cases)
def test_success_v0(arr: list[int], ans: bool, solution: Solution):
    assert ans == solution.uniqueOccurrences(arr)


@pytest.mark.parametrize(("arr", "ans"), test_cases)
def test_success_v1(arr: list[int], ans: bool, solution: Solution):
    assert ans == solution.uniqueOccurrencesV1(arr)


@pytest.fixture
def solution() -> Solution:
    return Solution()
