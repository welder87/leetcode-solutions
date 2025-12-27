# INTERVALS

1. Сначала отсортировть. 99% задач с интервалами начинаются с сортировки.
2. Выбор того, как сортировать: по start для слияния/поиска, по end для жадного выбора.
3. Держи "текущий" интервал и сравнивай его со следующим.

4. Для поиска "правого" интервала — бинарный поиск по массиву начал.
5. События (start/end) превращают интервалы в точки, что часто проще.
6. Проверка пересечения: max(start1, start2) < min(end1, end2).

## №1. Паттерны для сравнения интервалов

### Проверка на полное покрытие (для Remove Covered Intervals)

```python
def is_covered(a, b):
    # a покрывает b?
    return a[0] <= b[0] and a[1] >= b[1]
```

#### Проверка на пересечение

```python

def overlap(a, b):
    return not (a[1] <= b[0] or b[1] <= a[0])  # они НЕ НЕ пересекаются
    # ИЛИ
    return max(a[0], b[0]) < min(a[1], b[1])  # есть общая область
```

#### Расширение интервала (для Merge)

```python
new_interval = [min(a[0], b[0]), max(a[1], b[1])]
```

## №2. Правило сортировки "сначала по началу, потом по концу" (или наоборот)

Для большинства задач (Merge Intervals, Non-overlapping Intervals)

```python
intervals.sort(key=lambda x: (x[0], -x[1]))  # начало по возрастанию, конец по убыванию
```

или

```python
intervals.sort(key=lambda x: (x[0], x[1]))   # начало и конец по возрастанию
```

Для задач, где важен конец (чаще для жадных алгоритмов)

```python
intervals.sort(key=lambda x: x[1])  # сортировка по концу
```

Почему `-x[1]`? В задаче Remove Covered Intervals это критично:

```text
Интервалы после сортировки: [[1,4], [1,3], [1,2], [2,3], [3,4]]
С `x[1]` по убыванию:      [[1,4], [1,3], [1,2], [2,3], [3,4]]
Теперь интервал с большим концом идет первым — легко обнаружить покрывающие!
```

## №3. Паттерн "Текущий активный интервал" (для Merge/Overlap)

```python
# Базовый шаблон для Merge Intervals
result = []
for interval in sorted_intervals:
    # Если result пуст ИЛИ текущий интервал не пересекается с последним в result
    if not result or result[-1][1] < interval[0]:
        result.append(interval)
    else:
        # Пересекаются — расширяем последний интервал
        result[-1][1] = max(result[-1][1], interval[1])  # КРИТИЧЕСКАЯ СТРОКА!
```

Вариация для Non-overlapping Intervals:

```python
# Жадный подход — берём интервалы, которые раньше заканчиваются
count = 0
last_end = float('-inf')
for start, end in sorted(intervals, key=lambda x: x[1]):  # сортировка по КОНЦУ
    if start >= last_end:  # не пересекается
        last_end = end     # берём этот интервал
    else:
        count += 1         # удаляем/пропускаем этот интервал
```

## №4. Идиома "Поиск правого интервала" (для Find Right Interval)

```python
# Шаг 1: Сохраняем оригинальные индексы перед сортировкой
intervals_with_index = [(start, end, i) for i, (start, end) in enumerate(intervals)]
intervals_with_index.sort()  # сортируем по start

# Шаг 2: Бинарный поиск для каждого интервала
import bisect
starts = [interval[0] for interval in sorted_intervals]  # массив начал

result = []
for start, end in original_intervals:
    # Ищем первое start_i >= end текущего интервала
    idx = bisect.bisect_left(starts, end)  # КЛЮЧЕВОЙ МОМЕНТ: ищем по КОНЦУ текущего
    if idx < len(starts):
        result.append(sorted_intervals[idx][2])  # возвращаем оригинальный индекс
    else:
        result.append(-1)
```

## №5. Паттерн "Сканирующая линия" (Sweep Line) для подсчёта перекрытий

Идиома для задач типа "максимальное количество одновременных интервалов"

```python
events = []
for start, end in intervals:
    events.append((start, 1))   # начало интервала: +1
    events.append((end, -1))    # конец интервала: -1

events.sort()  # сортируем по времени, при равенстве сначала идут "-1" (концы)

current = 0
max_overlap = 0
for time, delta in events:
    current += delta
    max_overlap = max(max_overlap, current)
```

## №6. Паттерн "Два указателя для интервалов"

```python
# Когда нужно сравнить интервалы из двух разных списков
# (например, найти пересечения интервалов)
i, j = 0, 0
while i < len(A) and j < len(B):
    a_start, a_end = A[i]
    b_start, b_end = B[j]

    # Проверяем пересечение
    lo = max(a_start, b_start)
    hi = min(a_end, b_end)
    if lo <= hi:  # есть пересечение
        result.append([lo, hi])

    # Двигаем указатель того интервала, который раньше заканчивается
    if a_end < b_end:
        i += 1
    else:
        j += 1
```

## №7. Идиома "Интервалы в виде событий" для сложных условий

```python
# Для задач типа "Meeting Rooms II", но с дополнительными условиями
meetings = []
for start, end in intervals:
    meetings.append((start, 'start'))
    meetings.append((end - 0.1, 'end'))  # магия! чтобы при равенстве начала и конца
                                         # начало обрабатывалось раньше

meetings.sort()  # теперь сортировка по времени, а при равенстве
                 # 'start' будет раньше из-за -0.1
```

## №8. Шаблон для вставки нового интервала

```python
# Для задачи Insert Interval
def insert(intervals, new_interval):
    result = []
    i = 0
    n = len(intervals)

    # 1. Добавляем все интервалы, которые заканчиваются ДО начала нового
    while i < n and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1

    # 2. Мержим все пересекающиеся интервалы
    while i < n and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    result.append(new_interval)

    # 3. Добавляем оставшиеся
    while i < n:
        result.append(intervals[i])
        i += 1

    return result
```
