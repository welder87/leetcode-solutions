from unittest.mock import patch
import pytest

from lc_278_first_bad_version import Solution


test_cases = (
    # preset cases
    (5, 4),
    (1, 1),
    # common cases
    (5, 1),
    (5, 2),
    (5, 3),
    (5, 5),
    (4, 1),
    (4, 2),
    (4, 3),
    (4, 4),
    (2, 1),
    (2, 2),
    # corner cases
)
path = "lc_278_first_bad_version.isBadVersion"


@pytest.mark.parametrize(("n", "ans"), test_cases)
def test_success_v0(n: int, ans: int, responses, solution: Solution):
    with patch(path) as mock:
        mock.side_effect = responses
        assert solution.firstBadVersion(n) == ans


@pytest.mark.parametrize(("n", "ans"), test_cases)
def test_success_v1(n: int, ans: int, responses, solution: Solution):
    path = "lc_278_first_bad_version.isBadVersion"
    with patch(path) as mock:
        mock.side_effect = responses
        assert solution.firstBadVersionV1(n) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()


@pytest.fixture
def responses(n, ans):
    values = {val: True if val >= ans else False for val in range(1, n + 1)}

    def _side_effect(arg):
        return values[arg]

    return _side_effect
