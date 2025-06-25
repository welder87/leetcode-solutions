from functools import partial
from unittest.mock import patch

import pytest
from typing_extensions import Literal

from lc_374_guess_number_higher_or_lower.lc_374_guess_number_higher_or_lower import (
    Solution,
)

test_cases = (
    # preset cases
    (10, 6),
    (1, 1),
    (2, 1),
    # common cases
    (100, 51),
    (99, 50),
    (99, 98),
    (99, 2),
    # corner cases
    (99, 99),
    (99, 1),
    (2, 2),
)


@pytest.mark.parametrize(("num", "ans"), test_cases)
def test_success_v0(num: int, ans: bool, solution: Solution):
    with patch(
        "lc_374_guess_number_higher_or_lower.lc_374_guess_number_higher_or_lower.guess",
    ) as mock:
        mock.side_effect = partial(_guess, ans)
        assert ans is solution.guessNumber(num)


@pytest.mark.parametrize(("num", "ans"), test_cases)
def test_success_v1(num: int, ans: bool, solution: Solution):
    with patch(
        "lc_374_guess_number_higher_or_lower.lc_374_guess_number_higher_or_lower.guess",
    ) as mock:
        mock.side_effect = partial(_guess, ans)
        assert ans is solution.guessNumberV1(num)


@pytest.fixture
def solution() -> Solution:
    return Solution()


def _guess(pick: int, num: int) -> Literal[0, -1, 1]:
    if num == pick:
        return 0
    if num < pick:
        return 1
    return -1
