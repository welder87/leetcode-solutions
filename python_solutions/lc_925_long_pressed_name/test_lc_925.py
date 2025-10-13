import pytest

from lc_925_long_pressed_name.lc_925 import Solution

test_cases = (
    # preset cases
    pytest.param("alex", "aaleex", True, id="1"),
    pytest.param("saeed", "ssaaedd", False, id="2"),
    pytest.param("alex", "aaleexa", False, id="3"),
    # common cases
    pytest.param("abcde", "aabcccdee", True, id="100"),
    pytest.param("abcde", "aabcccdfee", False, id="101"),
    pytest.param("rnd", "ernd", False, id="102"),
    pytest.param("rnnnd", "rrnnddd", False, id="103"),
    pytest.param("rnnnd", "rrnnnddd", True, id="104"),
    pytest.param("abcde", "aabcccdeef", False, id="105"),
    # corner cases
    pytest.param("b", "b", True, id="200"),
    pytest.param("bc", "bc", True, id="201"),
    pytest.param("bc", "bcc", True, id="202"),
    pytest.param("bc", "bcd", False, id="203"),
    pytest.param("bc", "bcdddfd", False, id="204"),
    pytest.param("b", "c", False, id="205"),
)


@pytest.mark.parametrize(("name", "typed", "ans"), test_cases)
def test_success_v0(name: str, typed: str, ans: bool, solution: Solution):
    assert solution.isLongPressedName(name, typed) is ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
