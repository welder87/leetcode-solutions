from operator import le, lt


class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(n log n) + O(log n) + O(log n) + O(m).
        # Space complexity: O(n) + O(1) + O(1) + O(m).
        nums.sort()
        left = self._calc_boundary(nums, target, "left")
        right = self._calc_boundary(nums, target, "right")
        return list(range(left, right + 1))

    def _calc_boundary(
        self,
        nums: list[int],
        target: int,
        boundary: str,
    ):
        comparer = lt if boundary == "left" else le
        left, right = -1, len(nums)
        while right > left + 1:
            mid = left + (right - left) // 2
            if comparer(nums[mid], target):
                left = mid
            else:
                right = mid
        return right if boundary == "left" else left

    def targetIndicesV1(self, nums: list[int], target: int) -> list[int]:
        # Time complexity: O(2n) + O(m) + O(k).
        # Space complexity: O(m) + O(k).
        counter = [0] * (max(nums) + 1)
        isFound = False
        for num in nums:
            if not isFound and num == target:
                isFound = True
            counter[num] += 1
        if not isFound:
            return []
        cnt = 0
        for num in range(1, len(counter)):
            if num == target:
                break
            cnt += counter[num]
        return list(range(cnt, cnt + counter[target]))
