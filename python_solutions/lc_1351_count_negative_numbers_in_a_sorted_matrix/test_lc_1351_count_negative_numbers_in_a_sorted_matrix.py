import pytest

from lc_1351_count_negative_numbers_in_a_sorted_matrix import Solution


test_cases = (
    # preset cases
    ([[4, 3, 2, -1], [3, 2, 1, -1], [1, 1, -1, -2], [-1, -1, -2, -3]], 8),
    ([[3, 2], [1, 0]], 0),
    # common cases
    (
        [
            [4, 3, 1, 0],
            [4, 3, 0, -1],
            [0, -1, -2, -3],
            [-1, -1, -2, -3],
            [4, 3, -1, -3],
            [-1, -1, -1, -1],
            [-1, -1, -1, -2],
            [0, -1, -1, -1],
        ],
        21,
    ),
    # corner cases
    ([[-1], [1]], 1),
)


@pytest.mark.parametrize(("grid", "ans"), test_cases)
def test_success_v0(grid: list[list[int]], ans: int, solution: Solution):
    assert solution.countNegatives(grid) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
