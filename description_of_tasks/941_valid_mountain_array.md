# 941. Valid Mountain Array

Given an array of integers `arr`, return _`true` if and only if it is a valid mountain array_.

Recall that arr is a mountain array if and only if:

* `arr.length >= 3`
* There exists some `i` with `0 < i < arr.length - 1` such that:
  * `arr[0] < arr[1] < ... < arr[i - 1] < arr[i]`
  * `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

![hint_valid_mountain_array](img/941_valid_mountain_array_hint_valid_mountain_array.png)

**Example 1:**

> **Input:** arr = \[2,1\]
>
> **Output:** false

**Example 2:**

> **Input:** arr = \[3,5,5\]
>
> **Output:** false

**Example 3:**

> **Input:** arr = \[0,3,2,1\]
>
> **Output:** true

**Constraints:**

* `1 <= arr.length <= 104`
* `0 <= arr[i] <= 104`

## Topics

* `Array`

## Hints

It's very easy to keep track of a monotonically increasing or decreasing ordering of elements. You just need to be able to determine the start of the valley in the mountain and from that point onwards, it should be a valley i.e. no mini-hills after that. Use this information in regards to the values in the array and you will be able to come up with a straightforward solution.
