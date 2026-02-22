import pytest

from lc_3136_valid_word.lc_3136 import Solution

test_cases = (
    pytest.param("234Adas", True, id="Preset Case 1"),
    pytest.param("b3", False, id="Preset Case 2"),
    pytest.param("a3$e", False, id="Preset Case 3"),
    pytest.param("aya", True, id="Preset Case 4"),
    pytest.param("IS", False, id="Preset Case 5"),
    pytest.param("a1b", True, id="Exactly 3 characters (vowel, consonant, digit)"),
    pytest.param("a1e", False, id="Exactly 3 characters (only vowels and digit)"),
    pytest.param("b2c", False, id="Exactly 3 characters (only consonants and digits)"),
    pytest.param("ab", False, id="Length 2 - too short"),
    pytest.param("a", False, id="Length 1 - too short"),
    pytest.param(
        "abcdefghij1234567890",
        True,
        id="Maximum length (20 chars) (has vowels and consonants)",
    ),
    pytest.param("123456", False, id="Only digits - no vowels or consonants"),
    pytest.param("123@456", False, id="Digits with special char"),
    pytest.param("@#$", False, id="Special character only"),
    pytest.param("abc@123", False, id="Mixed valid and invalid chars"),
    pytest.param("BCDFG", False, id="All valid chars (no vowel)"),
    pytest.param("aeiou", False, id="All valid chars (no consonant)"),
    pytest.param("a123bc", True, id="Single lowercase vowel with consonants"),
    pytest.param("E123bc", True, id="Single uppercase vowel with consonants"),
    pytest.param("aeiouAEIOU", False, id="All vowels - no consonants"),
    pytest.param("aei123ou", False, id="Mixed vowels and digits - no consonants"),
    pytest.param("apple123", True, id="Vowel at beginning"),
    pytest.param("123bana", True, id="Vowel at end"),
    pytest.param("aeioub", True, id="All vowels except one consonant"),
    pytest.param("b123ae", True, id="Single consonant with vowels"),
    pytest.param("bcdfghjklmnpqrstvwxyz", False, id="All consonants - no vowels"),
    pytest.param("bcd123fgh", False, id="Consonants with digits - no vowels"),
    pytest.param("ball123", True, id="Consonant at beginning"),
    pytest.param("123aeiob", True, id="Consonant at end"),
    pytest.param("bcdfghja", True, id="All consonants except one vowel"),
    pytest.param("HeLlO123", True, id="Mixed case - valid"),
    pytest.param("AEIOUBCD123", True, id="Uppercase vowels"),
    pytest.param("BCDFG", False, id="Uppercase consonants only"),
    pytest.param("AEIOU", False, id="Uppercase vowels only"),
    pytest.param("aeiou", False, id="Lowercase vowels only"),
    pytest.param("abc123def", True, id="Mixed alphanumeric with both requirements"),
    pytest.param("supercalifragilisticexpialidocious123", True, id="Long valid word"),
    pytest.param("a1b2c3d4e5f6", True, id="Numbers interspersed"),
    pytest.param("123hello", True, id="Starts with number"),
    pytest.param("hello123", True, id="Ends with number"),
    pytest.param("bcdfghjklmnpa", True, id="Only one vowel among many consonants"),
    pytest.param("aeiouaeioub", True, id="Only one consonant among many vowels"),
    pytest.param("bcd123fgh", False, id="No vowels with digits"),
    pytest.param("aei123ou", False, id="No consonants with digits"),
    pytest.param("abc@123def", False, id="Special character in middle"),
    pytest.param("abc@123", False, id="Single allowed special char '@' - invalid"),
    pytest.param("abc#123", False, id="Single allowed special char '#' - invalid"),
    pytest.param("abc$123", False, id="Single allowed special char '$' - invalid"),
    pytest.param("123", False, id="Minimum length with digits only"),
    pytest.param("bcd", False, id="Minimum length with letters only (no vowel)"),
    pytest.param("aei", False, id="Minimum length with letters only (no consonant)"),
    pytest.param("ab1", True, id="Minimum length valid"),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: bool, solution: Solution):
    assert ans is solution.isValid(s)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v1(s: str, ans: bool, solution: Solution):
    assert ans is solution.isValidV1(s)


@pytest.fixture
def solution() -> Solution:
    return Solution()
