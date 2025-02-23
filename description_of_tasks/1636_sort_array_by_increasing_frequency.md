# 1636. Sort Array by Increasing Frequency

Given an array of integers `nums`, sort the array in **increasing** order based on the frequency of the values. If multiple values have the same frequency, sort them in **decreasing** order.

Return the _sorted array_.

**Example 1:**

> **Input:** nums = \[1,1,2,2,2,3\]
>
> **Output:** \[3,1,1,2,2,2\]
>
> **Explanation:** '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.

**Example 2:**

> **Input:** nums = \[2,3,1,3,2\]
>
> **Output:** \[1,3,3,2,2\]
>
> **Explanation:** '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.

**Example 3:**

> **Input:** nums = \[-1,1,-6,4,5,-6,1,4,1\]
>
> **Output:** \[5,-1,4,4,-6,-6,1,1,1\]

## Constraints

* `1 <= nums.length <= 100`
* `-100 <= nums[i] <= 100`

## Topics

* `Array`
* `Hash Table`
* `Sorting`

## Hints

1. Count the frequency of each value.
2. Use a custom comparator to compare values by their frequency. If two values have the same frequency, compare their values.

## Similar Questions

Easy

* [2206. Divide Array Into Equal Pairs]()
* [2190. Most Frequent Number Following Key In an Array]()
* [2341. Maximum Number of Pairs in Array]()
* [2418. Sort the People]()

Medium

* [451. Sort Characters By Frequency]()
* [2374. Node With Highest Edge Score]()

## Solution

### Overview

The task is to sort an array of integers by their frequency, placing numbers with fewer occurrences first. If two numbers appear with the same frequency, they should be ordered by their values in descending order. Think of it as arranging a playlist by the least to most popular songs, or ranking search results to prioritize the most relevant and engaging options for a more intuitive user experience.

* * *

### Approach: Customized Sorting

#### Intuition

To sort the numbers, we first arrange them based on their frequency in ascending order. Numbers that appear less frequently will come before those with higher frequencies. We use a hashmap,`freq`, to count the occurrences of each number in the array.

If two numbers have the same frequency, we then sort them by their values in descending order. This introduces a dual sorting criterion: first by frequency and then by value.

To accomplish this, we will apply a custom sorting function using lambda expressions. These anonymous functions let us define sorting logic inline. Specifically, our lambda function ensures that numbers are compared primarily by their frequency, and secondarily by their value if frequencies match. This approach guarantees that the final sorted list adheres to both sorting criteria.

##### Python Lambda Function for Sorting by Increasing Frequency

    sorted(nums, key=lambda x: (freq[x], -x))

The lambda function`lambda x: (freq[x], -x)`is used as the`key`parameter in the`sorted`function call.

1. `lambda x:`creates an anonymous function with`x`as its parameter.
2. `(freq[x], -x)`is the tuple that the lambda function returns.
3. `freq[x]`is used to get the frequency of`x`from the`freq`dictionary as the main sorting criterion.
4. `-x`ensures that values are sorted in descending order when their frequencies are the same.

#### Algorithm

* Initialize an unordered map`freq`to store the frequency of each integer in the input array`nums`.
* Traverse through each integer`num`in the array`nums`.
* Increase the count of`num`in the`freq`map using`freq[num]++`.
* Sort the array`nums`using the`sort`function with a custom comparator:
  * Compare two integers`a`and`b`based on their frequencies stored in the`freq`map:
    * If`freq[a]`(frequency of`a`) equals`freq[b]`(frequency of`b`), then:
    * Return`a > b`to ensure that in case of tie-in frequency, larger values come first (decreasing order).
    * Otherwise, return`freq[a] < freq[b]`to sort by frequency in increasing order.
* Return the sorted`nums`array, which now reflects the integers sorted primarily by frequency in ascending order, and by value in descending order when frequencies are tied.

#### Implementation

```python
from collections import Counter


class Solution:
    def frequencySort(self, nums: List[int]) -> List[int]:
        freq = Counter(nums)
        return sorted(nums, key=lambda x: (freq[x], -x))
```

#### Complexity Analysis

LetNbe the length of`nums`.

* Time complexity: $O(N\log N)$.

    Sorting`nums`incurs a time complexity of  $O(N\log N)$. Iterating over`nums`when counting frequencies incurs a time complexity of  $O(N)$, which can be ignored since $O(N\log N)$ is the dominating term.

* Space complexity: $O(N)$. We define a hash map to count the frequencies of each element, which incurs a space complexity of $O(N)$. Sorting also takes up some space, and the space complexity for that is detailed below:

    Some extra space is used when we sort an array of sizeNin place. The space complexity of the sorting algorithm depends on the programming language.

  * In Python, the`sort`method sorts a list using the Timsort algorithm which is a combination of Merge Sort and Insertion Sort and has a space complexity of $O(N)$

    Overall, the worst-case time complexity will be $O(N)$ when the array`nums`is filled with unique elements.
