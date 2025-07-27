import pytest

from common_datastructures import to_singly_linked_list
from lc_234_palindrome_linked_list.lc_234_palindrome_linked_list import Solution

test_cases = (
    # preset cases
    ([1, 2, 2, 1], True),
    ([1, 2], False),
    # common cases
    ([1, 2, 1], True),
    ([1, 2, 5, 2, 1], True),
    ([1, 2, 8, 1], False),
    ([1, 2, 2, 0], False),
    # corner cases
    ([0], True),
    ([1, 1], True),
    ([1, 1, 1], True),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: bool, solution: Solution):
    assert ans is solution.isPalindrome(to_singly_linked_list(lst))


@pytest.fixture
def solution() -> Solution:
    return Solution()
