import pytest

from lc_929_unique_email_addresses import Solution


test_cases = (
    (
        [
            "test.email+alex@leetcode.com",
            "test.e.mail+bob.cathy@leetcode.com",
            "testemail+david@lee.tcode.com",
        ],
        2,
    ),
    (["a@leetcode.com", "b@leetcode.com", "c@leetcode.com"], 3),
    (
        [
            "test.email+alex+@leet+code.com",
            "test.e.mail+bob.ca+thy@l.eet+code.com",
            "testemai.l+david@lee.tcode.com",
        ],
        3,
    ),
    (["test.email+alex@leetcode.com"], 1),
    ([], 0),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[str], ans: int, solution: Solution):
    assert ans is solution.numUniqueEmails(nums)


@pytest.fixture
def solution() -> Solution:
    return Solution()
