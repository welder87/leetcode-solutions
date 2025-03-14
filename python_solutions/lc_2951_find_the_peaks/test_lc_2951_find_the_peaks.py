import pytest

from lc_2951_find_the_peaks import Solution


test_cases = (
    # preset cases
    ([2, 4, 4], set()),
    ([1, 4, 3, 8, 5], {1, 3}),
    # common cases
    ([7, 6, 4, 5, 6, 7, 8], set()),
    ([7, 6, 5, 5, 8, 6, 3, 9, 8], {4, 7}),
    ([7, 6, 4, 3, 2, 1, 5], set()),
    ([1, 2, 3, 4, 4, 3, 2], set()),
    ([1, 2, 3, 4, 5, 4, 3, 2], {4}),
    # corner cases
    ([7, 6, 4], set()),
    ([7, 8, 4], {1}),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: set[int], solution: Solution):
    assert set(solution.findPeaks(nums)) == ans


@pytest.fixture
def solution() -> Solution:
    return Solution()
