import pytest

from lc_69_sqrt_x.lc_69_sqrt_x import Solution

test_cases = (
    # preset cases
    (4, 2),
    (8, 2),
    # common cases
    (1_000_000, 1000),
    (100, 10),
    (99, 9),
    (25, 5),
    (9, 3),
    (7, 2),
    (5, 2),
    (3, 1),
    (2, 1),
    # corner cases
    (1, 1),
)


@pytest.mark.parametrize(("num", "ans"), test_cases)
def test_success_v0(num: int, ans: int, solution: Solution):
    assert ans == solution.mySqrt(num)


@pytest.fixture
def solution() -> Solution:
    return Solution()
