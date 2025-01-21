class Solution:
    def leftRightDifference(self, nums: list[int]) -> list[int]:
        length = len(nums)
        left = [0] * length
        for i in range(length - 1):
            left[i + 1] = left[i] + nums[i]
        right = [0] * length
        for i in range(length - 1, 0, -1):
            right[i - 1] = right[i] + nums[i]
        result = [None] * length
        for idx, (elem1, elem2) in enumerate(zip(left, right)):
            result[idx] = abs(elem1 - elem2)
        return result

    def leftRightDifferenceV1(self, nums: list[int]) -> list[int]:
        length = len(nums)
        last_index = length - 1
        result = [0] * length
        for i in range(last_index):
            result[i + 1] = result[i] + nums[i]
        prev = 0
        for i in range(last_index, -1, -1):
            result[i] = abs(prev - result[i])
            prev += nums[i]
        return result
