# 485. Max Consecutive Ones

Given a binary array `nums`, return _the maximum number of consecutive_ `1`_'s in the array_.

**Example 1:**

> **Input:** nums = \[1,1,0,1,1,1\]
>
> **Output:** 3
>
> **Explanation:** The first two digits or the last three digits are > consecutive 1s. The maximum number of consecutive 1s is 3.

**Example 2:**

> **Input:** nums = \[1,0,1,1,0,1\]
>
> **Output:** 2

## Constraints

* `1 <= nums.length <= 105`
* `nums[i]` is either `0` or `1`.

## Topics

* `Array`

## Hints

1. You need to think about two things as far as any window is concerned. One is the starting point for the window. How do you detect that a new window of 1s has started? The next part is detecting the ending point for this window. How do you detect the ending point for an existing window? If you figure these two things out, you will be able to detect the windows of consecutive ones. All that remains afterward is to find the longest such window and return the size.
