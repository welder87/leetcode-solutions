import pytest

from lc_2129_capitalize_the_title.lc_2129 import Solution

test_cases = (
    pytest.param("capiTalIze tHe titLe", "Capitalize The Title", id="Preset Case 1"),
    pytest.param(
        "First leTTeR of EACH Word",
        "First Letter of Each Word",
        id="Preset Case 2",
    ),
    pytest.param("i lOve leetcode", "i Love Leetcode", id="Preset Case 3"),
    pytest.param("ZW Cl pyR uoC", "zw cl Pyr Uoc", id="Preset Case 4"),
    pytest.param("a", "a", id="Single character word"),
    pytest.param("an", "an", id="Single 2-character word"),
    pytest.param("the", "The", id="Single 3-character word"),
    pytest.param("word", "Word", id="4-character word should be capitalized"),
    pytest.param("A AN THE", "a an The", id="All uppercase short words"),
    pytest.param(
        "it is a test of and but for",
        "it is a Test of And But For",
        id="Mixed case with boundary length",
    ),
    pytest.param(
        "the quick brown fox jumps over the lazy dog",
        "The Quick Brown Fox Jumps Over The Lazy Dog",
        id="Long title with mixed lengths",
    ),
    pytest.param(
        f"a {'verylongword' * 10}",
        f"a V{('verylongword' * 10)[1:]}",
        id="Very long word",
    ),
    pytest.param(
        "THE CAT AND THE HAT",
        "The Cat And The Hat",
        id="All caps short words",
    ),
    pytest.param("PROGRAMMING IS FUN", "Programming is Fun", id="All caps long words"),
    pytest.param("an in on at to be", "an in on at to be", id="Exactly 2 characters"),
    pytest.param(
        "the and for but nor yet",
        "The And For But Nor Yet",
        id="Exactly 3 characters",
    ),
    pytest.param(
        "this that what when where",
        "This That What When Where",
        id="Exactly 4 characters",
    ),
)


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v0(s: str, ans: str, solution: Solution):
    assert solution.capitalizeTitle(s) == ans


@pytest.mark.parametrize(("s", "ans"), test_cases)
def test_success_v1(s: str, ans: str, solution: Solution):
    assert solution.capitalizeTitleV1(s) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
