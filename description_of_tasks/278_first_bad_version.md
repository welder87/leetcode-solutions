# 278. First Bad Version

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

**Example 1:**

> **Input:** n = 5, bad = 4
>
> **Output:** 4
>
> **Explanation:**
>
> call isBadVersion(3) -> false
>
> call isBadVersion(5) -> true
>
> call isBadVersion(4) -> true
>
> Then 4 is the first bad version.

**Example 2:**

> **Input:** n = 1, bad = 1
>
> **Output:** 1

## Constraints

* $1$ <= bad <= n <= $2^{31} - 1$

## Topics

* `Array`
* `Interactive`

## Similar Questions

Medium

* [34. Find First and Last Position of Element in Sorted Array](34_find_first_and_last_position_of_element_in_sorted_array.md)

Easy

* [35. Search Insert Position](35_search_insert_position.md)
* [374. Guess Number Higher or Lower]()
