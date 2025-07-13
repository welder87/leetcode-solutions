class Solution:
    def countPairs(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n*n). Space complexity: O(1).
        if len(nums) <= 1:
            return 0
        ans = 0
        for i in range(0, len(nums) - 1):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] < target:
                    ans += 1
        return ans

    def countPairsV1(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n log n) + O(n log n). Space complexity: O(n).
        # modified itmo
        if len(nums) <= 1:
            return 0
        nums.sort()
        ans = 0
        for i in range(len(nums) - 1):
            left, right = i + 1, len(nums)
            while left + 1 < right:
                mid = left + (right - left) // 2
                if i < mid and nums[i] + nums[mid] < target:
                    left = mid
                else:
                    right = mid
            if i < left and nums[i] + nums[left] < target:
                ans += left - i
        return ans

    def countPairsV2(self, nums: list[int], target: int) -> int:
        # Time complexity: O(n) + O(n)+ O(max N + 1) + O(n).
        # Space complexity: O((max N + 1).
        if len(nums) <= 1:
            return 0
        # counting sort
        min_num, max_num = nums[0], nums[0]
        for i in range(1, len(nums)):
            if min_num > nums[i]:
                min_num = nums[i]
            if max_num < nums[i]:
                max_num = nums[i]
        diff = 0
        if min_num < 0:
            diff = abs(min_num)
        sorted_nums = [0] * (abs(max_num) + diff + 1)
        for num in nums:
            sorted_nums[num + diff] += 1
        i = 0
        for num, cnt in enumerate(sorted_nums):
            if cnt == 0:
                continue
            val = num - diff
            for _ in range(cnt):
                nums[i] = val
                i += 1
        sorted_nums.clear()
        sorted_nums = None
        # two pointers
        ans = 0
        i, j = 0, len(nums) - 1
        while i < j:
            if nums[i] + nums[j] < target:
                ans += j - i
                i += 1
            else:
                j -= 1
        return ans
