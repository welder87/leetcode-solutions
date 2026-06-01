# Singly Linked List

## Микропаттерны

### Dummy Node

Ключевая идея: создаем dummy, работаем с указателями, а в конце возвращаем
`dummy->next`.

#### Применимость

Dummy node полезен, когда:

- Удаление головы списка.
    - [19](19.md)
    - [82](82.md)
    - [203](203.md)
- Вставка в начало списка.
    - [707](707.md)
    - [1472](1472.md)
- Строим новый список и не знаем заранее его голову.
    - [2](2.md)
    - [21](21.md)
    - [86](86.md)
    - [206](206.md)
    - [328](328.md)
- Нужен "предыдущий" указатель для узла, который может быть головой.
    - [19](19.md)
- Поиск или удаление N-го узла списка от начала.
    - [876](876.md)
    - [2095](2095.md)

## Паттерны

### Two pointers (Обобщенный)

Примеры задач

- [19](19.md)
- [21](21.md)

### Merge sorted array

Примеры задач

- [21](21.md)

### Fast, slow

Быстрый указатель сдвигаем на 2 узла.
Медленный указатель сдвигаем на 1 узел.
Условие выхода могут различаться в зависимости от задачи.

```go
fast, slow := head, head
for condition {
    fast = fast.Next.Next
    slow = slow.Next
}
```

Примеры задач

- [141](141.md)
- [876](876.md)

### Reverse

Примеры задач

- [2487](2487.md)
- [2816](2816.md)

### Iteration

Итерация по списку.

```go
for node != nil {
    node = node.Next
}
```

### Length

Расчет длины списка.

Без Dummy Node.

```go
length := 0
for node != nil {
    length++
    node = node.Next
}
```

C Dummy Node.

```go
length := 0
node := dummy
for node != nil {
    node = node.Next
    length++
}
```

Примеры задач

- [19](19.md)

### Removal from the beginning

```go

```

Примеры задач

- [2181](2181.md)
- [3217](3217.md)

## Примеры задач

```go
// Объявление структуры односвязного списка
type ListNode struct {
     Val int
     Next *ListNode
}
```

### 2. Add Two Numbers

#### [Description](https://leetcode.com/problems/add-two-numbers/)

#### Solution

```go

```

### 23. Merge k Sorted Lists

#### [Description](https://leetcode.com/problems/merge-k-sorted-lists)

#### Solution

```go

```

### 86. Partition List

#### [Description](https://leetcode.com/problems/partition-list/)

#### Solution

```go

```

### 203. Remove Linked List Elements

#### [Description](https://leetcode.com/problems/remove-linked-list-elements/)

#### Solution

```go

```

### 707. Design Linked List

#### [Description](https://leetcode.com/problems/design-linked-list)

#### Solution

```go

```

### 1472. Design Browser History

#### [Description](https://leetcode.com/problems/design-browser-history)

#### Solution

```go

```
