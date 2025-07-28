import pytest

from common_datastructures import to_singly_linked_list
from lc_234_palindrome_linked_list.lc_234_palindrome_linked_list import Solution

test_cases = (
    # preset cases
    pytest.param([1, 2, 2, 1], True, id="0"),
    pytest.param([1, 2], False, id="2"),
    # common cases
    pytest.param([1, 2, 1], True, id="3"),
    pytest.param([1, 2, 5, 2, 1], True, id="4"),
    pytest.param([1, 2, 8, 1], False, id="5"),
    pytest.param([1, 2, 2, 0], False, id="6"),
    pytest.param([1, 3, 2, 6, 7], False, id="7"),
    pytest.param([1, 9, 0, 8], False, id="8"),
    # corner cases
    pytest.param([0], True, id="9"),
    pytest.param([1, 1], True, id="10"),
    pytest.param([1, 1, 1], True, id="11"),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: bool, solution: Solution):
    assert ans is solution.isPalindrome(to_singly_linked_list(lst))


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v1(lst: list[int], ans: bool, solution: Solution):
    assert ans is solution.isPalindromeV1(to_singly_linked_list(lst))


@pytest.fixture
def solution() -> Solution:
    return Solution()
