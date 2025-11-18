import pytest

from lc_283_move_zeroes.lc_283 import Solution

test_cases = (
    # PRESET CASES
    pytest.param([0, 1, 0, 3, 12], [1, 3, 12, 0, 0], id="Preset case 1"),
    pytest.param([0], [0], id="Preset case 2"),
    # COMMON CASES
    # все ноли в конце. массив не меняется.
    pytest.param([1, 3, 12, 0, 0], [1, 3, 12, 0, 0], id="Trailing zeros"),
    # все ноли в начале
    pytest.param([0, 0, 0, 1, 2, 3], [1, 2, 3, 0, 0, 0], id="Leading zeros"),
    # ноли в середине. порядок ненулевых элементов
    pytest.param([1, 0, 2, 0, 3, 0], [1, 2, 3, 0, 0, 0], id="Embedded zeros"),
    # чередующиеся ноли и не ноли. обработка чередующегося паттерна - старт с ноля.
    pytest.param(
        [0, 1, 0, 2, 0, 3, 0, 4],
        [1, 2, 3, 4, 0, 0, 0, 0],
        id="Alternating zeros and non-zeros (zero start)",
    ),
    # чередующиеся ноли и не ноли. обработка чередующегося паттерна - старт с не ноля.
    pytest.param(
        [1, 0, 2, 0, 3, 0, 4],
        [1, 2, 3, 4, 0, 0, 0],
        id="Alternating zeros and non-zeros (non-zero start)",
    ),
    # ноли в начале. ноли в конце.
    pytest.param(
        [0, 0, 0, 1, 3, 12, 0, 0],
        [1, 3, 12, 0, 0, 0, 0, 0],
        id="Leading and trailing zeros",
    ),
    # большие числа
    pytest.param(
        [1000000, 0, 2**31, 0, 0, 42],
        [1000000, 2**31, 42, 0, 0, 0],
        id="Large numbers",
    ),
    # CORNER CASES
    # нет нолей. массив не меняется.
    pytest.param([1, 6, 1, 24], [1, 6, 1, 24], id="No zeros"),
    # все элементы нули
    pytest.param([0, 0, 0], [0, 0, 0], id="All zeros"),
    # один элемент - ноль
    pytest.param([0], [0], id="Single Zero"),
    # один элемент - не ноль
    pytest.param([5], [5], id="Single Non-Zero"),
)


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v0(nums: list[int], ans: set[int], solution: Solution):
    nums_id = id(nums)
    solution.moveZeroes(nums)
    assert nums == ans
    assert id(nums) == nums_id


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v1(nums: list[int], ans: set[int], solution: Solution):
    nums_id = id(nums)
    solution.moveZeroesV1(nums)
    assert nums == ans
    assert id(nums) == nums_id


@pytest.mark.parametrize(("nums", "ans"), test_cases)
def test_success_v2(nums: list[int], ans: set[int], solution: Solution):
    nums_id = id(nums)
    solution.moveZeroesV2(nums)
    assert nums == ans
    assert nums == ans
    assert id(nums) == nums_id


@pytest.fixture
def solution() -> Solution:
    return Solution()
