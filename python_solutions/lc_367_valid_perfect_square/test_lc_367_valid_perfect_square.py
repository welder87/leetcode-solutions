import pytest

from lc_367_valid_perfect_square.lc_367_valid_perfect_square import Solution

test_cases = (
    # preset cases
    (16, True),
    (14, False),
    # common cases
    (1_000_000, True),
    (100, True),
    (99, False),
    (25, True),
    (9, True),
    (7, False),
    (4, True),
    (3, False),
    (2, False),
    # corner cases
    (1, True),
)


@pytest.mark.parametrize(("num", "ans"), test_cases)
def test_success_v0(num: int, ans: bool, solution: Solution):
    assert ans is solution.isPerfectSquare(num)


@pytest.fixture
def solution() -> Solution:
    return Solution()
