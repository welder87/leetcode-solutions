import pytest

from lc_415_add_strings import Solution



test_cases = (
    ("11", "123", "134"),
    ("456", "77", "533"),
    ("0", "0", "0"),
    ("1", "9", "10"),
    ("1", "999", "1000"),
    ("268", "342", "610"),
    ("768", "343", "1111"),
)


@pytest.mark.parametrize(("nums1", "nums2", "ans"), test_cases)
def test_success_v0(nums1: str, nums2: str, ans: str, solution: Solution):
    assert ans == solution.addStrings(nums1, nums2)


@pytest.fixture
def solution() -> Solution:
    return Solution()
