import pytest

from lc_383_ransom_note.lc_383 import Solution

test_cases = (
    pytest.param("a", "b", False, id="Preset case 1"),
    pytest.param("aa", "ab", False, id="Preset case 2"),
    pytest.param("ab", "aab", True, id="Preset case 3"),
    pytest.param("a", "ab", True, id="Valid construction - simple"),
    pytest.param("aa", "aab", True, id="Valid construction - exact match"),
    pytest.param("aa", "ab", False, id="Invalid - insufficient letters"),
    pytest.param("abc", "ab", False, id="Invalid - missing letter"),
    pytest.param("hello", "helloworld", True, id="Valid - magazine has extra letters"),
    pytest.param("abc", "cba", True, id="Valid - same letters different order"),
    pytest.param("aaa", "aa", False, id="Invalid - missing one instance"),
    pytest.param(
        "aaaaaaaaaa",
        "aaaaaaaaaaaaaaaaaaaa",
        True,
        id="Large ransom note valid",
    ),
    pytest.param("aaaaaaaaaa", "aaaaaaaaa", False, id="Large ransom note invalid"),
    pytest.param("aaaaab", "aaaaac", False, id="Multiple same character needed"),
    pytest.param(
        "aabbcc",
        "aaabbbccc",
        True,
        id="Multiple characters with correct frequency",
    ),
)


@pytest.mark.parametrize(("ransom_note", "magazine", "ans"), test_cases)
def test_success_v0(ransom_note: str, magazine: str, ans: bool, solution: Solution):
    assert ans is solution.canConstruct(ransom_note, magazine)


@pytest.mark.parametrize(("ransom_note", "magazine", "ans"), test_cases)
def test_success_v1(ransom_note: str, magazine: str, ans: bool, solution: Solution):
    assert ans is solution.canConstructV1(ransom_note, magazine)


@pytest.fixture
def solution() -> Solution:
    return Solution()
