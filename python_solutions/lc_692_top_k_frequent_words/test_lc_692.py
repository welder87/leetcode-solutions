import pytest

from lc_692_top_k_frequent_words.lc_692 import Solution

test_cases = (
    pytest.param(
        ["i", "love", "leetcode", "i", "love", "coding"],
        2,
        ["i", "love"],
        id="preset_case_1",
    ),
    pytest.param(
        ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
        4,
        ["the", "is", "sunny", "day"],
        id="preset_case_2",
    ),
    pytest.param(
        ["i", "blow", "leetcode", "i", "blow", "coding"],
        2,
        ["blow", "i"],
        id="long_word_vs_short_word__with_same_frequent",
    ),
    pytest.param(
        ["i", "k", "e", "i", "k", "e"],
        3,
        ["e", "i", "k"],
        id="three_words_with_single_character_1",
    ),
    pytest.param(
        ["i", "k", "e", "i", "k", "e"],
        2,
        ["e", "i"],
        id="three_words_with_single_character_2",
    ),
    pytest.param(
        ["i", "k", "e", "i", "k", "e"],
        1,
        ["e"],
        id="three_words_with_single_character_3",
    ),
    pytest.param(["i"], 1, ["i"], id="single_word_with_single_character"),
    pytest.param(["item"], 1, ["item"], id="single_word"),
    pytest.param(["a", "a", "a"], 1, ["a"], id="all words same"),
    pytest.param(
        ["apple", "banana", "apple", "banana", "cherry"],
        2,
        ["apple", "banana"],
        id="same frequency, different lexicographical order",
    ),
    pytest.param(
        ["dog", "cat", "bat", "rat"],
        4,
        ["bat", "cat", "dog", "rat"],
        id="multiple_words_with_same_frequency",
    ),
    pytest.param(
        ["zoo", "zoo", "apple", "banana", "banana"],
        3,
        ["banana", "zoo", "apple"],
        id="mix_of_frequencies_and_lexicographical_order",
    ),
    pytest.param(
        ["word"] * 1000 + ["another"] * 500 + ["third"] * 250,
        3,
        ["word", "another", "third"],
        id="performance_case",
    ),
)


@pytest.mark.parametrize(("words", "k", "ans"), test_cases)
def test_success_v0(words: list[str], k: int, ans: list[str], solution: Solution):
    assert solution.topKFrequent(words, k) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
