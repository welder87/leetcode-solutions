import pytest

from lc_986_interval_list_intersections.lc_986 import Solution

test_cases = (
    pytest.param(
        [[0, 2], [5, 10], [13, 23], [24, 25]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        [[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]],
        id="Preset case 1",
    ),
    pytest.param([[1, 3], [5, 9]], [], [], id="Preset case 2"),
    pytest.param(
        [[3, 5], [9, 20]],
        [[4, 5], [7, 10], [11, 12], [14, 15], [16, 20]],
        [[4, 5], [9, 10], [11, 12], [14, 15], [16, 20]],
        id="complex case",
    ),
    pytest.param(
        [[0, 30]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        id="nested within one (first)",
    ),
    pytest.param(
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        [[0, 30]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        id="nested within one (second)",
    ),
    pytest.param([[1, 10]], [[3, 7]], [[3, 7]], id="one inside another (first)"),
    pytest.param([[3, 7]], [[1, 10]], [[3, 7]], id="one inside another (second)"),
    pytest.param(
        [[1, 4], [6, 8]],
        [[2, 4], [5, 8]],
        [[2, 4], [6, 8]],
        id="ending at same point",
    ),
    pytest.param(
        [[1, 3], [5, 7]],
        [[1, 2], [5, 6]],
        [[1, 2], [5, 6]],
        id="starting at same point",
    ),
    pytest.param(
        [[0, 2], [5, 10]],
        [[10, 24], [15, 24]],
        [[10, 10]],
        id="touch at point (first-second)",
    ),
    pytest.param(
        [[10, 24], [15, 24]],
        [[0, 2], [5, 10]],
        [[10, 10]],
        id="touch at point (second-first)",
    ),
    pytest.param(
        [[1, 3], [4, 6]],
        [[3, 4]],
        [[3, 3], [4, 4]],
        id="touch at point (middle)",
    ),
    pytest.param(
        [[24, 25]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        [[24, 24], [25, 25]],
        id="Partial overlap of two (first)",
    ),
    pytest.param(
        [[1, 5], [8, 12], [15, 24], [25, 26]],
        [[24, 25]],
        [[24, 24], [25, 25]],
        id="Partial overlap of two (second)",
    ),
    pytest.param([[1, 4]], [[3, 6]], [[3, 4]], id="partial overlap right"),
    pytest.param([[3, 6]], [[1, 4]], [[3, 4]], id="partial overlap left"),
    pytest.param(
        [[1, 3], [7, 9]],
        [[4, 6], [10, 12]],
        [],
        id="non-overlapping intervals",
    ),
    pytest.param(
        [[0, 2], [5, 10]],
        [[15, 24], [15, 24]],
        [],
        id="non-overlapping intervals (first)",
    ),
    pytest.param(
        [[15, 24], [15, 24]],
        [[0, 2], [5, 10]],
        [],
        id="non-overlapping intervals (second)",
    ),
    pytest.param([[0, 0]], [[1, 1]], [], id="disjoint point"),
    pytest.param(
        [[1, 3], [4, 6]],
        [[1, 3], [4, 6]],
        [[1, 3], [4, 6]],
        id="identical intervals",
    ),
    pytest.param([], [[1, 3], [5, 9]], [], id="empty first"),
    pytest.param([[1, 3], [5, 9]], [], [], id="empty second"),
    pytest.param([], [], [], id="both empty"),
)


@pytest.mark.parametrize(("first", "second", "ans"), test_cases)
def test_success_v0(
    first: list[list[int]],
    second: list[list[int]],
    ans: list[list[int]],
    solution: Solution,
):
    assert solution.intervalIntersection(first, second) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
