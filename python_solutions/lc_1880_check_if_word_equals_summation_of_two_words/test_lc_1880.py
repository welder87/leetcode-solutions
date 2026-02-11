import pytest

from helpers import to_pytest_id
from lc_1880_check_if_word_equals_summation_of_two_words.lc_1880 import Solution

test_cases = (
    pytest.param("acb", "cba", "cdb", True, id=to_pytest_id("Preset Case 1")),
    pytest.param("aaa", "a", "aab", False, id=to_pytest_id("Preset Case 2")),
    pytest.param("aaa", "a", "aaaa", True, id=to_pytest_id("Preset Case 3")),
    pytest.param("a", "a", "b", False, id=to_pytest_id("Single letters all a")),
    pytest.param("a", "a", "a", True, id=to_pytest_id("Single letters all a")),
    pytest.param("a", "b", "b", True, id=to_pytest_id("Single letters correct sum")),
    pytest.param(
        "a",
        "b",
        "c",
        False,
        id=to_pytest_id("Single letters incorrect sum"),
    ),
    pytest.param("a", "a", "a", True, id=to_pytest_id("All zero letters")),
    pytest.param("aaa", "a", "aaaa", True, id=to_pytest_id("Zero sum")),
    pytest.param("aab", "c", "d", True, id=to_pytest_id("Leading zeros in first word")),
    pytest.param(
        "d",
        "aac",
        "f",
        True,
        id=to_pytest_id("Leading zeros in second word"),
    ),
    pytest.param(
        "aad",
        "aae",
        "aai",
        False,
        id=to_pytest_id("Leading zeros in both words"),
    ),
    pytest.param("j" * 8, "a", "j" * 8, True, id=to_pytest_id("long words")),
    pytest.param(
        "jjjjjjji",
        "b",
        "jjjjjjjj",
        True,
        id=to_pytest_id("biggest target word"),
    ),
    pytest.param(
        "b",
        "jjjjjjji",
        "jjjjjjjj",
        True,
        id=to_pytest_id("small first big second"),
    ),
)


@pytest.mark.parametrize(("first", "second", "target", "ans"), test_cases)
def test_success_v0(first: str, second: str, target: str, ans: int, solution: Solution):
    assert solution.isSumEqual(first, second, target) is ans


@pytest.mark.parametrize(("first", "second", "target", "ans"), test_cases)
def test_success_v1(first: str, second: str, target: str, ans: int, solution: Solution):
    assert solution.isSumEqual(first, second, target) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
