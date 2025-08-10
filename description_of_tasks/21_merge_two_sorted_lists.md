# 21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return _the head of the merged linked list_.

**Example 1:**

![1](img/21_merge_two_sorted_lists.jpg)

> **Input:** list1 = [1,2,4], list2 = [1,3,4]
>
> **Output:** [1,1,2,3,4,4]

**Example 2:**

> **Input:** list1 = [], list2 = []
>
> **Output:** []

**Example 3:**

> **Input:** list1 = [], list2 = [0]
>
> **Output:** [0]

## Constraints

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

## Topics

- `Linked List`
- `Recursion`

## Similar Questions

Hard

- [23. Merge k Sorted Lists](23_merge_k_sorted_lists.md)

Medium

- [148. Sort List](148_sort_list.md)

Easy

- [88. Merge Sorted Array](88_merge_sorted_array.md)
- [2570. Merge Two 2D Arrays by Summing Values](2570_merge_two_2_d_arrays_by_summing_values.md)
