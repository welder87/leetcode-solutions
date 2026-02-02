import pytest

from lc_290_word_pattern.lc_290 import Solution

test_cases = (
    pytest.param("abba", "dog cat cat dog", True, id="preset_case_1"),
    pytest.param("abba", "dog cat cat fish", False, id="preset_case_2"),
    pytest.param("aaaa", "dog cat cat dog", False, id="preset_case_3"),
    pytest.param("abba", "dog dog dog dog", False, id="preset_case_4"),
    pytest.param("aaaa", "dog dog dog dog", True, id="preset_case_5"),
    pytest.param("abcaa", "dog cat fish cat cat", False, id="Error in end"),
    pytest.param("a", "dog cat cat dog", False, id="length_mismatch_1"),
    pytest.param("abba", "dog", False, id="length_mismatch_2"),
    pytest.param("abc", "dog cat", False, id="length_mismatch_3"),
    pytest.param("ab", "word", False, id="length_mismatch_5"),
    pytest.param("a", "dog", True, id="single_character_pattern_1"),
    pytest.param("a", "d", True, id="single_character_pattern_2"),
    pytest.param("abcaa", "dog cat fish dog dogs", False, id="Similar words"),
    pytest.param(
        "abc",
        "hello world there",
        True,
        id="same_pattern__different_words_1",
    ),
    pytest.param(
        "abc",
        "hello world hello",
        False,
        id="same_pattern__different_words_2",
    ),
    pytest.param(
        "abc",
        "hello world there",
        True,
        id="same_words__different_patterns_1",
    ),
    pytest.param(
        "abc",
        "hello world hello",
        False,
        id="same_words__different_patterns_2",
    ),
    pytest.param(
        "jquery",
        "jquery",
        False,
        id="Single word but pattern has multiple chars",
    ),
    pytest.param("ab", "dog dog", False, id="different_chars_cannot_map_to_same_word"),
    pytest.param("aa", "dog cat", False, id="same_char_cannot_map_to_different_words"),
    pytest.param("ab", "cat cat", False, id="both_issues"),
    pytest.param("abcdefghij", "a b c d e f g h i j", True, id="longer_pattern_1"),
    pytest.param(
        "aaaaaaaaaa",
        "word word word word word word word word word word",
        True,
        id="longer_pattern_2",
    ),
    pytest.param("abcabc", "cat dog fish cat dog fish", True, id="longer_pattern_3"),
    pytest.param("abcabc", "cat dog fish cat dog bird", False, id="longer_pattern_4"),
    pytest.param("a" * 1000, "word " * 999 + "word", True, id="performance_test_1"),
    pytest.param(
        "ab" * 500, ("word1 word2 " * 500).strip(), True, id="performance_test_1"
    ),
    pytest.param(
        "a" * 1000,
        " ".join([f"word{i}" for i in range(1000)]),
        False,
        id="performance_test_4",
    ),
)


@pytest.mark.parametrize(("pattern", "s", "ans"), test_cases)
def test_success_v0(pattern: str, s: str, ans: bool, solution: Solution):
    assert solution.wordPattern(pattern, s) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
