import pytest

from lc_1475_final_prices_with_a_spec_discount_in_a_shop.lc_1475 import Solution

test_cases = (
    # preset cases
    ([8, 4, 6, 2, 3], [4, 2, 4, 2, 3]),
    ([1, 2, 3, 4, 5], [1, 2, 3, 4, 5]),
    ([10, 1, 1, 6], [9, 0, 1, 6]),
    ([8, 7, 4, 2, 8, 1, 7, 7, 10, 1], [1, 3, 2, 1, 7, 0, 0, 6, 9, 1]),
    # common cases
    ([8, 7, 4, 2, 8, 1, 7, 8, 7, 1], [1, 3, 2, 1, 7, 0, 0, 1, 6, 1]),
    ([8, 9, 10, 3], [5, 6, 7, 3]),
    ([5, 4, 3, 1], [1, 1, 2, 1]),
    # corner cases
    ([10], [10]),
    ([10, 20], [10, 20]),
    ([10, 10], [0, 10]),
    ([10, 1], [9, 1]),
    ([1, 1, 1, 1], [0, 0, 0, 1]),
)


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v0(lst: list[int], ans: list[int], solution: Solution):
    assert solution.finalPrices(lst) == ans


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v1(lst: list[int], ans: list[int], solution: Solution):
    assert solution.finalPricesV1(lst) == ans


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v2(lst: list[int], ans: list[int], solution: Solution):
    assert solution.finalPricesV2(lst) == ans


@pytest.mark.parametrize(("lst", "ans"), test_cases)
def test_success_v3(lst: list[int], ans: list[int], solution: Solution):
    assert solution.finalPricesV3(lst) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
