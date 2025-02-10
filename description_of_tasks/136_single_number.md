# 136. Single Number

Given a **non-empty** array of integers `nums`, every element appears _twice_ except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

**Example 1:**

> **Input:** nums = \[2,2,1\]
>
> **Output:** 1

**Example 2:**

> **Input:** nums = \[4,1,2,1,2\]
>
> **Output:** 4

**Example 3:**

> **Input:** nums = \[1\]
>
> **Output:** 1

## Constraints

* `1 <= nums.length <= 3 * 104`
* `-3 * 104 <= nums[i] <= 3 * 104`
* Each element in the array appears twice except for one element which appears only once.

## Topics

* `Array`
* `Bit Manipulation`

## Hints

1. Think about the XOR (^) operator's property.

## Similar Questions

Easy

* [268. Missing Number](268_missing_number.md)
* [389. Find the Difference]()
* [3158. Find the XOR of Numbers Which Appear Twice]()

Medium

* [137. Single Number II]()
* [260. Single Number III]()
* [287. Find the Duplicate Number](287_find_the_duplicate_number.md)
