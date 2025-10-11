import pytest

from lc_821_shortest_distance_to_a_character.lc_821 import Solution

test_cases = (
    # preset cases
    pytest.param("loveleetcode", "e", [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0], id="1"),
    pytest.param("aaab", "b", [3, 2, 1, 0], id="2"),
    # common cases
    pytest.param(
        "abceabceabeabcd",
        "e",
        [3, 2, 1, 0, 1, 2, 1, 0, 1, 1, 0, 1, 2, 3, 4],
        id="100",
    ),
    pytest.param("abceabc", "e", [3, 2, 1, 0, 1, 2, 3], id="101"),
    pytest.param("eabc", "e", [0, 1, 2, 3], id="102"),
    pytest.param("ceabc", "e", [1, 0, 1, 2, 3], id="103"),
    pytest.param("rceabc", "e", [2, 1, 0, 1, 2, 3], id="104"),
    pytest.param("eceabe", "e", [0, 1, 0, 1, 1, 0], id="105"),
    # corner cases
    pytest.param("bbb", "b", [0, 0, 0], id="200"),
    pytest.param("bab", "a", [1, 0, 1], id="201"),
    pytest.param("bab", "b", [0, 1, 0], id="202"),
    pytest.param("b", "b", [0], id="203"),
    pytest.param("bc", "b", [0, 1], id="204"),
    pytest.param("bc", "c", [1, 0], id="205"),
)


@pytest.mark.parametrize(("s", "c", "ans"), test_cases)
def test_success_v0(s: str, c: str, ans: list[int], solution: Solution):
    assert solution.shortestToChar(s, c) == ans


@pytest.mark.parametrize(("s", "c", "ans"), test_cases)
def test_success_v1(s: str, c: str, ans: list[int], solution: Solution):
    assert solution.shortestToCharV1(s, c) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
