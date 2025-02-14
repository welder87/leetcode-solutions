713. Subarray Product Less Than K

Given an array of integers `nums` and an integer `k`, return _the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than_ `k`.

**Example 1:**

> **Input:** nums = \[10,5,2,6\], k = 100
>
> **Output:** 8
>
> **Explanation:** The 8 subarrays that have product less than 100 are:
>
> \[10\], \[5\], \[2\], \[6\], \[10, 5\], \[5, 2\], \[2, 6\], \[5, 2, 6\]
>
> Note that \[10, 5, 2\] is not included as the product of 100 is not strictly less than k.

**Example 2:**

> **Input:** nums = \[1,2,3\], k = 0
>
> **Output:** 0

## Constraints

* `1 <= nums.length <= 3 * 104`
* `1 <= nums[i] <= 1000`
* `0 <= k <= 106`

## Topics

* `Array`
* `Binary Search`
* `Sliding Window`
* `Prefix Sum`
* `Prefix Sum`
* `Two Pointers`

## Hints

1. For each j, let opt(j) be the smallest i so that nums\[i\] \* nums\[i+1\] \* ... \* nums\[j\] is less than k. opt is an increasing function.

* * *

## Similar Questions

Easy

* [Two Sum Less Than K]()

Medium

* [Maximum Product Subarray]()
* [Maximum Size Subarray Sum Equals k]()
* [Subarray Sum Equals K]()
* [Number of Smooth Descent Periods of a Stock]()

Hard

* [Count Subarrays With Score Less Than K]()
