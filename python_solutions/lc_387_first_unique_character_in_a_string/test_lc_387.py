import pytest

from helpers import to_pytest_id
from lc_387_first_unique_character_in_a_string.lc_387 import Solution

test_cases = (
    pytest.param("leetcode", 0, id=to_pytest_id("Preset Case 1")),
    pytest.param("loveleetcode", 2, id=to_pytest_id("Preset Case 2")),
    pytest.param("aabb", -1, id=to_pytest_id("Preset Case 3")),
    pytest.param("a", 0, id=to_pytest_id("Single character")),
    pytest.param("aa", -1, id=to_pytest_id("Two same characters")),
    pytest.param("ab", 0, id=to_pytest_id("Two different characters")),
    pytest.param("aaaaa", -1, id=to_pytest_id("All same character")),
    pytest.param("abcde", 0, id=to_pytest_id("First character is unique")),
    pytest.param("aabbc", 4, id=to_pytest_id("Last character is unique")),
    pytest.param("aabcc", 2, id=to_pytest_id("Middle character is unique")),
    pytest.param(
        "abc",
        0,
        id=to_pytest_id("Multiple uniques, first one should be returned"),
    ),
    pytest.param(
        "abcdefghijklmnopqrstuvwxyz",
        0,
        id=to_pytest_id("Long string with early unique"),
    ),
    pytest.param(
        "aaaaaaaaaaaaaaaaaaaaab",
        21,
        id=to_pytest_id("Long string with late unique"),
    ),
    pytest.param("racecar", 3, id=to_pytest_id("Palindrome with unique middle")),
    pytest.param("abccba", -1, id=to_pytest_id("Palindrome no unique")),
    pytest.param("abcabcabcx", 9, id=to_pytest_id("Repeating pattern with one unique")),
    pytest.param("aaabbbcccddde", 12, id=to_pytest_id("Many duplicates then unique")),
    pytest.param("aabbcbbaa", 4, id=to_pytest_id("Unique between duplicates")),
    pytest.param("ababab" * 100, -1, id=to_pytest_id("Performance Case")),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: int, solution: Solution):
    assert ans == solution.firstUniqChar(s)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v1(s: str, ans: int, solution: Solution):
    assert ans == solution.firstUniqCharV1(s)


@pytest.fixture
def solution() -> Solution:
    return Solution()
