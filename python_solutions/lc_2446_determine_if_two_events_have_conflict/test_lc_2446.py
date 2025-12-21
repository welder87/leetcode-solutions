import pytest

from lc_2446_determine_if_two_events_have_conflict.lc_2446 import Solution

test_cases = (
    # preset cases
    # |-----|
    #       |-----|
    (["01:15", "02:00"], ["02:00", "03:00"], True),
    # |-----|
    #     |-----|
    (["01:00", "02:00"], ["01:20", "03:00"], True),
    # |-----|
    #         |-----|
    (["10:00", "11:00"], ["14:00", "15:00"], False),
    # |-------|
    #   |--|
    (["01:37", "14:20"], ["05:06", "06:17"], True),
    # common cases
    #      |-----|
    # |-----|
    (["01:20", "03:00"], ["01:00", "02:00"], True),
    # |-----|
    #       |-----|
    (["01:15", "01:17"], ["01:17", "02:00"], True),
    (["21:15", "21:17"], ["21:17", "23:00"], True),
    #       |-----|
    # |-----|
    (["01:17", "02:00"], ["01:15", "01:17"], True),
    (["21:17", "23:00"], ["21:15", "21:17"], True),
    # |-----|
    # |-----|
    (["01:15", "02:00"], ["01:15", "02:00"], True),
    # |-----|
    #         |-----|
    (["23:15", "23:16"], ["23:00", "23:14"], False),
    #         |-----|
    # |-----|
    (["23:00", "23:14"], ["23:15", "23:16"], False),
    #    |--|
    # |-------|
    (["14:15", "14:16"], ["01:00", "23:55"], True),
    # |-------|
    #   |--|
    (["01:00", "23:55"], ["14:15", "14:16"], True),
)


@pytest.mark.parametrize(("event1", "event2", "ans"), test_cases)
def test_success_v0(
    event1: list[str],
    event2: list[str],
    ans: bool,
    solution: Solution,
):
    assert ans == solution.haveConflict(event1, event2)


@pytest.mark.parametrize(("event1", "event2", "ans"), test_cases)
def test_success_v1(
    event1: list[str],
    event2: list[str],
    ans: bool,
    solution: Solution,
):
    assert ans == solution.haveConflictV1(event1, event2)


@pytest.fixture
def solution() -> Solution:
    return Solution()
