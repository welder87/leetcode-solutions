import pytest

from lc_3110_score_of_a_string.lc_3110 import Solution

test_cases = (
    pytest.param("hello", 13, id="Preset Case 1"),
    pytest.param("zaz", 50, id="Preset Case 2"),
    pytest.param("aa", 0, id="Two identical characters"),
    pytest.param("ab", 1, id="Two consecutive alphabet characters"),
    pytest.param("ab", 1, id="Reverse consecutive"),
    pytest.param("aaaaa", 0, id="All same character"),
    pytest.param("zzzzzzzz", 0, id="All same character long"),
    pytest.param("abcde", 4, id="Increasing sequence"),
    pytest.param("abcdefghijklmnopqrstuvwxyz", 25, id="Full alphabet increasing"),
    pytest.param("edcba", 4, id="Decreasing sequence"),
    pytest.param("zyxwvutsrqponmlkjihgfedcba", 25, id="Full alphabet decreasing"),
    pytest.param("azaza", 100, id="Alternating high-low"),
    pytest.param("zazaz", 100, id="Alternating low-high"),
    pytest.param("az", 25, id="Maximum difference"),
    pytest.param("za", 25, id="Maximum to minimum"),
    pytest.param("ababababab", 9, id="Long repeating pattern"),
    pytest.param("abcba", 4, id="Palindrome odd"),
    pytest.param("abccba", 4, id="Palindrome even"),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: int, solution: Solution):
    assert solution.scoreOfString(s) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
