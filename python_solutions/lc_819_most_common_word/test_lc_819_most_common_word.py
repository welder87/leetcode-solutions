import pytest

from lc_819_most_common_word import Solution


test_cases = (
    # preset cases
    ("Bob hit a ball, the hit BALL flew far after it was hit.", ["hit"], "ball"),
    ("a.", [], "a"),
    # common cases
    (
        "Bob hit a ball, the hit BALL flew far ,after it, was. hit flew.",
        ["hit", "ball"],
        "flew",
    ),
    (
        "Bob hit a ball, the hit BALL flew far ,after it, was. hit flew.",
        ["hit", "ball", "flew"],
        "bob",
    ),
    # corner cases
    ("a", ["a"], ""),
    ("a", ["v"], "a"),
)


@pytest.mark.parametrize(("paragraph", "banned", "ans"), test_cases)
def test_success_v0(paragraph: str, banned: list[str], ans: str, solution: Solution):
    assert ans == solution.mostCommonWord(paragraph, banned)


@pytest.fixture
def solution() -> Solution:
    return Solution()
