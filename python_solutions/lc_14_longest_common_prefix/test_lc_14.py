import pytest

from lc_14_longest_common_prefix.lc_14 import Solution

test_cases = (
    pytest.param(["flower", "flow", "flight"], "fl", id="Preset Case 1"),
    pytest.param(["dog", "racecar", "car"], "", id="Preset Case 2"),
    pytest.param(["flower"], "flower", id="Single non empty word"),
    pytest.param([""], "", id="Single empty word"),
    pytest.param(["", ""], "", id="Multiple empty strings"),
    pytest.param(["a"], "a", id="Single character"),
    pytest.param(["a", "a"], "a", id="Same single characters"),
    pytest.param(["a", "b"], "", id="Different single characters"),
    pytest.param(["abc", "abc", "abc"], "abc", id="All identical"),
    pytest.param(["flower", "flow", "fright"], "f", id="Single character in end"),
    pytest.param(["fright", "flower", "flow"], "f", id="Single character in start"),
    pytest.param(["flower", "flowe", "flow"], "flow", id="Full word in end"),
    pytest.param(["flow", "flower", "flowe"], "flow", id="Full word in start"),
    pytest.param(["flow", "flower", ""], "", id="Empty word in end"),
    pytest.param(["", "flower", "flowe"], "", id="Empty word in start"),
    pytest.param(["throne", "throne"], "throne", id="Two same words"),
    pytest.param(
        ["a" * 1000, "a" * 1000, "a" * 1000],
        "a" * 1000,
        id="Performance Case 1",
    ),
    pytest.param(
        [f"pre{'x' * 1000}", f"pre{'y' * 1000}"],
        "pre",
        id="Performance Case 2",
    ),
    pytest.param(
        ["abcdefgh" * 100, "abcdefgh" * 100],
        "abcdefgh" * 100,
        id="Performance Case 3",
    ),
    pytest.param(
        ["interspecies", "interstellar", "interstate"],
        "inters",
        id="Success Case 1",
    ),
)


@pytest.mark.parametrize(("strs", "ans"), test_cases)
def test_success_v0(strs: list[str], ans: str, solution: Solution):
    assert ans == solution.longestCommonPrefix(strs)


@pytest.mark.parametrize(("strs", "ans"), test_cases)
def test_success_v1(strs: list[str], ans: str, solution: Solution):
    assert ans == solution.longestCommonPrefix(strs)


@pytest.fixture
def solution() -> Solution:
    return Solution()
