import pytest

from lc_1422_maximum_score_after_splitting_a_string import Solution


test_cases = (
    # preset cases
    ("011101", 5),
    ("00111", 5),
    ("1111", 3),
    # common cases
    ("011110", 5),
    ("10101", 3),
    ("111110", 4),
    ("011111", 6),
    ("000111", 6),
    # corner cases
    ("0000", 3),
    ("00", 1),
    ("10", 0),
    ("01", 2),
    ("11", 1),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: int, solution: Solution):
    assert ans == solution.maxScore(s)


@pytest.fixture
def solution() -> Solution:
    return Solution()
