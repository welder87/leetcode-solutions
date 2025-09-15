import pytest

from lc_581_shortest_unsorted_continuous_subarray.lc_581 import Solution

test_cases = (
    # preset cases
    pytest.param([2, 6, 4, 8, 10, 9, 15], 5, id="1"),
    pytest.param([1, 2, 3, 4], 0, id="2"),
    pytest.param([1], 0, id="3"),
    # common cases
    pytest.param([2, 6, 4, 8, 10, 9, 15, 25, 34], 5, id="100"),
    pytest.param([2, 6, 4, 8, 5, 9, 25, 3, 34], 7, id="101"),
    pytest.param([2, 2, 2, 1, 3], 4, id="102"),
    pytest.param([4, 1, 2, 3], 4, id="103"),
    pytest.param([2, 2, 2], 0, id="104"),
    pytest.param([3, 3, 3, 2, 2, 2], 6, id="105"),
    pytest.param([3, 3, 3, 2, 2, 2, 3], 6, id="106"),
    pytest.param([-3, -3, -3, 2, 2, 2, 3], 0, id="107"),
    pytest.param([-3, 2, -3, 7, 1, -2, 0], 6, id="108"),
    pytest.param([-3, 2, -3, 7, 1, -2, 15], 5, id="109"),
    pytest.param([3, 2, -3, 7, 1, -2, 17], 6, id="110"),
    pytest.param([3, 2, 3, 3, 7, 8], 2, id="111"),
    pytest.param([-1, -1, 3, 0, 7, 8], 2, id="112"),
    pytest.param([-1, -1, 0, 3, 8, 7], 2, id="113"),
    pytest.param([-1, -1, 0, 3, -1, 7], 3, id="114"),
    pytest.param([4, 3, 2, 1], 4, id="115"),
    # corner cases
    pytest.param([1, 2], 0, id="200"),
    pytest.param([1, 0], 2, id="201"),
    pytest.param([1, -1], 2, id="202"),
    pytest.param([0, -1, 1], 2, id="203"),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(
    nums: list[int],
    ans: list[int],
    solution: Solution,
):
    res = solution.findUnsortedSubarray(nums)
    assert res == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
