import pytest

from lc_3794_reverse_string_prefix.lc_3794 import Solution

test_cases = (
    pytest.param("abcd", 2, "bacd", id="Preset Case 1"),
    pytest.param("xyz", 3, "zyx", id="Preset Case 2"),
    pytest.param("hey", 1, "hey", id="Preset Case 2"),
    pytest.param("abcdef", 6, "fedcba", id="Full reverse (even)"),
    pytest.param("abcde", 5, "edcba", id="Full reverse (odd)"),
    pytest.param("abcdef", 3, "cbadef", id="First three (even)"),
    pytest.param("abcde", 3, "cbade", id="First three (odd)"),
    pytest.param("abcdef", 4, "dcbaef", id="First four (even)"),
    pytest.param("abcde", 4, "dcbae", id="First four (odd)"),
    pytest.param("a", 1, "a", id="One char"),
    pytest.param("aa", 2, "aa", id="Same chars"),
)


@pytest.mark.parametrize(("s", "k", "ans"), test_cases)
def test_success_v0(s: str, k: int, ans: str, solution: Solution):
    assert ans == solution.reversePrefix(s, k)


@pytest.fixture
def solution() -> Solution:
    return Solution()
