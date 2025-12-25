import pytest

from lc_56_merge_intervals.lc_56 import Solution

test_cases = (
    pytest.param(
        [[1, 3], [2, 6], [8, 10], [15, 18]],
        [[1, 6], [8, 10], [15, 18]],
        id="Preset case 1 (overlapping intervals)",
    ),
    pytest.param(
        [[1, 4], [4, 5]],
        [[1, 5]],
        id="Preset case 2 (adjacent intervals)",
    ),
    pytest.param(
        [[4, 7], [1, 4]],
        [[1, 7]],
        id="Preset case 3 (unsorted overlapping)",
    ),
    pytest.param(
        [[1, 1], [2, 3], [2, 6], [2, 15]],
        [[1, 1], [2, 15]],
        id="Sorted overlapping",
    ),
    pytest.param(
        [[4, 7]],
        [[4, 7]],
        id="Single interval",
    ),
    pytest.param(
        [[1, 2], [3, 6], [7, 15]],
        [[1, 2], [3, 6], [7, 15]],
        id="No overlaps",
    ),
    pytest.param(
        [[1, 4], [2, 3], [0, 5]],
        [[0, 5]],
        id="All overlapping (first)",
    ),
    pytest.param(
        [[1, 3], [2, 6], [5, 15], [15, 18]],
        [[1, 18]],
        id="All overlapping (second)",
    ),
    pytest.param(
        [[1, 4], [3, 5], [2, 4], [7, 10]],
        [[1, 5], [7, 10]],
        id="Unsorted nested intervals (first)",
    ),
    pytest.param(
        [[7, 15], [1, 1], [1, 6], [15, 15]],
        [[1, 6], [7, 15]],
        id="Unsorted nested intervals (second)",
    ),
    pytest.param(
        [[1, 2], [2, 3], [3, 4]],
        [[1, 4]],
        id="Touching intervals multiple",
    ),
    pytest.param(
        [[2, 3], [2, 2], [3, 3], [1, 3], [5, 7], [2, 2], [4, 6]],
        [[1, 3], [4, 7]],
        id="Unsorted with gaps",
    ),
)


@pytest.mark.parametrize(("intervals", "ans"), test_cases)
def test_success_v0(
    intervals: list[list[int]],
    ans: list[list[int]],
    solution: Solution,
):
    assert solution.merge(intervals) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
