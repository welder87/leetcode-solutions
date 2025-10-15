class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        # Time complexity: O(n + n). Space complexity: O(1).
        i, j = 0, 1
        cnt = 1
        counter = 0
        # first pass through the array
        # marking as empty value
        while i < len(nums) or j < len(nums):
            while j < len(nums) and nums[i] == nums[j]:
                cnt += 1
                if cnt > 2:
                    nums[j] = None
                    counter += 1
                j += 1
            i = j
            j += 1
            cnt = 1
        j = 1
        # second pass through the array
        # moving empty values
        for i in range(len(nums)):
            if nums[i] is not None:
                j += 1
                continue
            while j < len(nums) and nums[j] is None:
                j += 1
            if j < len(nums):
                nums[i], nums[j] = nums[j], nums[i]
        return len(nums) - counter

    def removeDuplicatesV1(self, nums: list[int]) -> int:
        # Time complexity: O(n). Space complexity: O(1).
        # Solution: https://algo.monster/liteproblems/80
        # Pointer to track the position for the next valid element
        write_index = 0
        # Iterate through each element in the array
        for current_num in nums:
            # Allow element if:
            # 1. We have less than 2 elements (write_index < 2), OR
            # 2. Current element is different from the element 2 positions back
            if write_index < 2 or current_num != nums[write_index - 2]:
                # Place the current element at the write position
                nums[write_index] = current_num
                # Move write pointer forward
                write_index += 1
        # Return the length of the modified array
        return write_index
