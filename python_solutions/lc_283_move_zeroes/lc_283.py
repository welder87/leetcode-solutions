class Solution:
    def moveZeroes(self, nums: list[int]) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        if len(nums) <= 1:
            return
        i, j = 0, 1
        while i < len(nums) and j < len(nums):
            if nums[i] == nums[j] == 0:
                j += 1
            elif nums[i] == 0 and nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            else:
                i += 1
                j += 1

    def moveZeroesV1(self, nums: list[int]) -> None:
        # Time complexity: O(2n). Space complexity: O(n).
        non_zero_nums = [None] * len(nums)
        i = 0
        for num in nums:
            if num != 0:
                non_zero_nums[i] = num
                i += 1
        for i, num in enumerate(non_zero_nums):
            nums[i] = 0 if num is None else num

    def moveZeroesV2(self, nums: list[int]) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        i, j = 0, 1
        while j < len(nums):
            if nums[i] == 0 and nums[j] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
            elif nums[i] != 0:
                i += 1
            j += 1

    def moveZeroesV3(self, nums: list[int]) -> None:
        # Time complexity: O(3n). Space complexity: O(n).
        # Solution: https://www.geeksforgeeks.org/problems/move-all-zeroes-to-end-of-array0751/1
        n = len(nums)
        temp = [0] * n
        # to keep track of the index in temp[]
        j = 0
        # Copy non-zero elements to temp[]
        for i in range(n):
            if nums[i] != 0:
                temp[j] = nums[i]
                j += 1
        # Fill remaining positions in temp[] with zeros
        while j < n:
            temp[j] = 0
            j += 1
        # Copy all the elements from temp[] to arr[]
        for i in range(n):
            nums[i] = temp[i]

    def moveZeroesV4(self, nums: list[int]) -> None:
        # Time complexity: O(2n). Space complexity: O(1).
        # Solution: https://www.geeksforgeeks.org/problems/move-all-zeroes-to-end-of-array0751/1
        # Count of non-zero elements
        count = 0
        # If the element is non-zero, replace the element at
        # index 'count' with this element and increment count
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[count] = nums[i]
                count += 1
        # Now all non-zero elements have been shifted to
        # the front. Make all elements 0 from count to end.
        while count < len(nums):
            nums[count] = 0
            count += 1

    def moveZeroesV5(self, nums: list[int]) -> None:
        # Time complexity: O(n). Space complexity: O(1).
        # Solution: https://www.geeksforgeeks.org/problems/move-all-zeroes-to-end-of-array0751/1
        # Solution: https://leetcode.com/problems/move-zeroes/editorial/
        # Pointer to track the position
        # for next non-zero element
        count = 0
        for i in range(len(nums)):
            # If the current element is non-zero
            if nums[i] != 0:
                # Swap the current element with
                # the 0 at index 'count'
                nums[i], nums[count] = nums[count], nums[i]
                # Move 'count' pointer to the next position
                count += 1
