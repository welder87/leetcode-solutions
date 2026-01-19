import pytest

from lc_3083_existence_of_a_substring_in_a_string_reverse.lc_3083 import Solution

test_cases = (
    pytest.param("leetcode", True, id="Preset case 1"),
    pytest.param("abcba", True, id="Preset case 2"),
    pytest.param("abcd", False, id="Preset case 3"),
    pytest.param("ausoee", True, id="Preset case 4"),
    pytest.param("a", False, id="Single character"),
    pytest.param("abcde", False, id="No matching pairs"),
    pytest.param("aa", True, id="Two same characters"),
    pytest.param("aaa", True, id="Appears multiple times"),
    pytest.param("aba", True, id="Palindrome odd length"),
    pytest.param("abba", True, id="Palindrome even length"),
    pytest.param("ababa", True, id="Palindrome"),
    pytest.param("aab", True, id="Single pair in start"),
    pytest.param("baa", True, id="Single pair in end"),
    pytest.param("abcdefghijklmnopqrstuvwxyz", False, id="Long string no repeats"),
    pytest.param("abcdefedcba", True, id="Long Palindrome"),
    pytest.param("abcabc", False, id="First or second"),
    pytest.param("xyyx", True, id="Pair in middle"),
    pytest.param("abccba", True, id="Multiple valid pair"),
    pytest.param("abca", False, id=" No length-2 substring in reverse"),
    pytest.param("a" * 100, True, id="All same character"),
    pytest.param("ab" * 50, True, id="Pattern ab"),
    pytest.param("abc" * 33 + "ab", False, id="Pattern with valid substring"),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: bool, solution: Solution):
    assert solution.isSubstringPresent(s) is ans


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v1(s: str, ans: bool, solution: Solution):
    assert solution.isSubstringPresentV1(s) is ans


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v2(s: str, ans: bool, solution: Solution):
    assert solution.isSubstringPresentV2(s) is ans


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v3(s: str, ans: bool, solution: Solution):
    assert solution.isSubstringPresentV3(s) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
